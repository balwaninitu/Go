//Package products is one of the subsystems of this application which is used to
//capture product master data records for this application
//This pkg is made up of several functions to help maintain products records by
//allowing creation/display/modification/deletion
package product

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"golivejd/sharvar"
	"golivejd/users"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//BookDet, a struct to capture booking data records
type ProdDet struct {
	ProdID      string
	Cat         string
	Shortdesc   string
	Sku         string
	Uprice      float64
	Pcur        string
	Availqty    int
	Minstockqty int
	Createdt    string
	Chgdt       string
	Deldt       string

	Errors map[string]string
}

//IDErrMsg is data structure to collect data entry error msgs for the purpose of data sanity
type IDErrMsg struct {
	ProdID string

	Errors map[string]string
}

//ProdData is map that maintains all product details

type ProdData map[string]*ProdDet

//
//ProdHandler is a handle to opened database connection
type ProdHandler struct {
	db *sql.DB
}

//NewProdHandler helps to call product functions using the open db handle
func NewProdHandler(db *sql.DB) *ProdHandler {
	return &ProdHandler{
		db: db,
	}
}

//func ProdFormValidate is used to validate product details during data entry
func (msg *ProdDet) ProdFormValidate() bool {

	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.ProdID) == "" {
		msg.Errors["ProdID"] = "Please enter Product Code"
	}

	if strings.TrimSpace(msg.Cat) == "" {
		msg.Errors["Cat"] = "Please enter Product Category"
	}

	if strings.TrimSpace(msg.Shortdesc) == "" {
		msg.Errors["Shortdesc"] = "Please enter Prodct Desription"
	}

	if strings.TrimSpace(msg.Sku) == "" {
		msg.Errors["Sku"] = "Please enter Unit of Measure (Kg/Liter/Meter)"
	}

	if msg.Uprice == 0 {
		msg.Errors["Price"] = "Please enter unit price (SGD)"
	}

	return len(msg.Errors) == 0
}

//func ValidateID will be used to validate product ID during booking display/updation/deletion
func (msg *IDErrMsg) ValidateID() bool {

	msg.Errors = make(map[string]string)

	sharvar.Mu.Lock()
	//get from database

	_, tRows := GetProdRecDB(msg.ProdID)

	sharvar.Mu.Unlock()

	if strings.TrimSpace(msg.ProdID) == "" {
		msg.Errors["ProdID"] = "Please enter Product Code"

	} else if tRows == 0 {

		msg.Errors["ProdID"] = "Product does not exist"
	}

	return len(msg.Errors) == 0
}

//func AddProd is used to add new product record into the system
func (h1 *ProdHandler) AddProd(res http.ResponseWriter, req *http.Request) {

	if !users.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if err := h1.db.Ping(); err != nil {
		panic(err)
	}

	if req.Method == http.MethodPost {

		//do string to number conversion

		//	Uprice1 := req.FormValue("Uprice")
		Uprice, _ := strconv.ParseFloat(req.FormValue("Uprice"), 64)
		Availqty, _ := strconv.Atoi(req.FormValue("Availqty"))
		Minstockqty, _ := strconv.Atoi(req.FormValue("Minstockqty  "))

		msg := &ProdDet{
			ProdID: req.FormValue("ProdID"),

			Cat:         req.FormValue("Cat"),
			Shortdesc:   req.FormValue("Shortdesc"),
			Sku:         req.FormValue("Sku"),
			Uprice:      Uprice,
			Availqty:    Availqty,
			Minstockqty: Minstockqty,

			Createdt: req.FormValue("Createdt"),
		}

		if msg.ProdFormValidate() == false {
			sharvar.Tpl.ExecuteTemplate(res, "prodadd.gohtml", msg)
			return
		}

		sharvar.Mu.Lock()

		//get rec from database
		_, tRows := GetProdRecDB(msg.ProdID)

		sharvar.Mu.Unlock()

		if tRows == 0 {

			sharvar.Mu.Lock()

			//store into database table
			_ = h1.AddProdRecDB(*msg)

			sharvar.Mu.Unlock()

			tmp := "<br>" + "<b>" + "ProdID" + " ( " + req.FormValue("ProdID") + " ) " + "added" + "</b>" + "<br>"
			_, _ = fmt.Fprintln(res, tmp)

		} else {

			tmp := "<br>" + "<b>" + " ( " + req.FormValue("ProdID") + " ) " + " already exists - can not add" + "</b>" + "<br>"
			_, _ = fmt.Fprintln(res, tmp)

		}

	}

	sharvar.Tpl.ExecuteTemplate(res, "prodadd.gohtml", nil)

}

//func DisProd will display details for a given product code
func (h1 *ProdHandler) DisProd(res http.ResponseWriter, req *http.Request) {

	_ = users.GetUser(res, req)
	if !users.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if err := h1.db.Ping(); err != nil {
		panic(err)
	}

	var (
		ProdID string
		tmp    string
	)

	if req.Method == http.MethodPost {

		msg := &IDErrMsg{
			ProdID: req.FormValue("ProdID"),
		}

		if msg.ValidateID() == false {
			sharvar.Tpl.ExecuteTemplate(res, "getprodIDWeb.gohtml", msg)
			return
		}

		ProdID = msg.ProdID
	}

	if ProdID != "" {

		sharvar.Mu.Lock()

		//get rec from database and copy the details onto screen fields
		prodTabRec, tRows := GetProdRecDB(ProdID)
		tmpProdDet := &prodTabRec

		sharvar.Mu.Unlock()

		if tRows > 0 {

			sharvar.Tpl.ExecuteTemplate(res, "proddis.gohtml", tmpProdDet)

		} else {

			tmp = "<br>" + "<b>" + "ProdID" + " ( " + req.FormValue("ProdID") + " ) " + " does not exists - can not display" + "</b>" + "<br>"
			_, _ = fmt.Fprintln(res, tmp)
		}
	} else {

		sharvar.Tpl.ExecuteTemplate(res, "getprodIDWeb.gohtml", nil)

	}
}

//func ModProd & func ModProdRec are used to update product details in the system
func (h1 *ProdHandler) ModProd(res http.ResponseWriter, req *http.Request) {

	_ = users.GetUser(res, req)
	if !users.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if err := h1.db.Ping(); err != nil {
		panic(err)
	}

	var (
		ProdID string
	)

	msg := &IDErrMsg{}
	tmpProdDet := &ProdDet{}

	if req.Method == http.MethodPost {

		msg = &IDErrMsg{
			ProdID: req.FormValue("ProdID"),
		}

		if msg.ValidateID() == false {
			sharvar.Tpl.ExecuteTemplate(res, "getprodIDWeb.gohtml", msg)
			return
		}

	}

	ProdID = msg.ProdID

	sharvar.Mu.Lock()
	//get rec from database and copy the details for screen fields
	prodTabRec, tRows := GetProdRecDB(ProdID)
	tmpProdDet = &prodTabRec

	sharvar.Mu.Unlock()

	if tRows > 0 {

		sharvar.Tpl.ExecuteTemplate(res, "prodmod.gohtml", tmpProdDet)

	} else {

		sharvar.Tpl.ExecuteTemplate(res, "getprodIDWeb.gohtml", nil)

	}

}

//func ModProdRec & func ModProd are used to update product details in the system
func (h1 *ProdHandler) ModProdRec(res http.ResponseWriter, req *http.Request) {

	if err := h1.db.Ping(); err != nil {
		panic(err)
	}

	var (
		ProdID    string
		Cat       string
		Shortdesc string
		Sku       string
		Uprice    float64
		Pcur      string
	//	Createdt  string
	//	Chgdt     string
	//	Deldt     string
	)

	ProdID = req.PostFormValue("ProdID")

	Cat = req.PostFormValue("Cat")
	Shortdesc = req.PostFormValue("Shortdesc")
	Sku = req.PostFormValue("Sku")
	//	Uprice = req.PostFormValue("Uprice")
	Pcur = req.PostFormValue("Pcur")

	//do string to number conversion

	//	Price1 = req.FormValue("Price")
	Uprice, _ = strconv.ParseFloat(req.FormValue("Uprice"), 64)

	msg := ProdDet{
		ProdID:    ProdID,
		Cat:       Cat,
		Shortdesc: Shortdesc,
		Sku:       Sku,
		Uprice:    Uprice,
		Pcur:      Pcur,
	}

	if msg.ProdFormValidate() == false {
		sharvar.Tpl.ExecuteTemplate(res, "prodmod.gohtml", msg)
		//		return

	} else {

		sharvar.Mu.Lock()
		//save to database
		h1.ModProdRecDB(msg)

		sharvar.Mu.Unlock()
		tmp := "<br>" + "<b>" + "ProdID" + " ( " + req.FormValue("ProdID") + " ) " + "modified" + "</b>" + "<br>"
		_, _ = fmt.Fprintln(res, tmp)
	}

	myUser := users.GetUser(res, req)
	sharvar.Tpl.ExecuteTemplate(res, "index.gohtml", myUser)

}

//func DelProduct & func DelProdRec  are used to delete product details
func (h1 *ProdHandler) DelProd(res http.ResponseWriter, req *http.Request) {

	_ = users.GetUser(res, req)
	if !users.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if err := h1.db.Ping(); err != nil {
		panic(err)
	}

	var ProdID string

	if req.Method == http.MethodPost {

		msg := &IDErrMsg{
			ProdID: req.FormValue("ProdID"),
		}

		if msg.ValidateID() == false {
			sharvar.Tpl.ExecuteTemplate(res, "getprodIDWeb.gohtml", msg)
			return
		}

		//	BookID = req.FormValue("BookID")
		ProdID = msg.ProdID
	}

	if ProdID != "" {

		//get rec from database and copy the details into screen fields
		prodTabRec, tRows := GetProdRecDB(ProdID)
		tmpProdDet := &prodTabRec

		if tRows > 0 {

			sharvar.Tpl.ExecuteTemplate(res, "proddel.gohtml", tmpProdDet)

		} else {

			tmp := "<br>" + "<b>" + "ProdID" + " ( " + req.FormValue("ProdID") + " ) " + " does not exists - can not delete" + "</b>" + "<br>"
			_, _ = fmt.Fprintln(res, tmp)
		}
	} else {

		sharvar.Tpl.ExecuteTemplate(res, "getprodIDWeb.gohtml", nil)
	}

}

//func DelProdRec func & DelProd are used to delete product details for a given booking ID
func (h1 *ProdHandler) DelProdRec(res http.ResponseWriter, req *http.Request) {

	if err := h1.db.Ping(); err != nil {
		panic(err)
	}

	ProdID := req.FormValue("ProdID")

	sharvar.Mu.Lock()

	//delete rec from database
	_ = h1.DelProdRecDB(ProdID)

	sharvar.Mu.Unlock()

	tmp := "<br>" + "<b>" + "ProdID" + " ( " + req.FormValue("ProdID") + " ) " + " deleted" + "</b>" + "<br>"
	_, _ = fmt.Fprintln(res, tmp)

	sharvar.Tpl.ExecuteTemplate(res, "proddel.gohtml", ProdID)

	myUser := users.GetUser(res, req)
	sharvar.Tpl.ExecuteTemplate(res, "index.gohtml", myUser)

}

//UplProdMast - upload product master
func (h1 *ProdHandler) UplProdMast(res http.ResponseWriter, req *http.Request) {

	// _ = GetUser(res, req)
	// if !AlreadyLoggedIn(req) {
	// 	http.Redirect(res, req, "/", http.StatusSeeOther)
	// 	return
	// }

	file := "C:\\Projects\\Go\\src\\golivejd\\assets\\prodmast.csv"

	var (
		msg       = ProdDet{}
		UpdStatus string
	)

	type resLine struct {
		UpdStatus   string
		ProdID      string
		Cat         string
		Shortdesc   string
		Sku         string
		Uprice      float64
		Pcur        string
		Availqty    int
		Minstockqty int
		Createdt    string
		Chgdt       string
		Deldt       string
	}
	var resLineDis = []resLine{}

	records, err := readProdRecs(file)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {

		Uprice1, _ := strconv.ParseFloat(record[4], 64)
		Availqty1, _ := strconv.Atoi(record[5])
		Minstockqty1, _ := strconv.Atoi(record[6])

		msg = ProdDet{
			ProdID:      strings.TrimSpace(record[0]),
			Cat:         record[1],
			Shortdesc:   record[2],
			Sku:         record[3],
			Uprice:      Uprice1,
			Availqty:    Availqty1,
			Minstockqty: Minstockqty1,
		}

		tmpProdDet, tRows := GetProdRecDB(msg.ProdID)

		if tRows == 0 {

			h1.AddProdRecDB(msg)
			UpdStatus = "New Rec Uploaded"

		} else {

			if tmpProdDet.Cat != msg.Cat || tmpProdDet.Shortdesc != msg.Shortdesc || tmpProdDet.Sku != msg.Sku || tmpProdDet.Uprice != msg.Uprice || tmpProdDet.Availqty != msg.Availqty || tmpProdDet.Minstockqty != msg.Minstockqty {

				//add to existing stock
				//		msg.Availqty = tmpProdDet.Availqty + msg.Availqty
				//		msg.Minstockqty = tmpProdDet.Minstockqty + msg.Minstockqty

				h1.ModProdRecDB(msg)
				UpdStatus = "Existing Rec Updated"
			} else {

				UpdStatus = "Rec No Changes Made"
			}

		}

		resLine1 := resLine{

			UpdStatus:   UpdStatus,
			ProdID:      strings.TrimSpace(record[0]),
			Cat:         record[1],
			Shortdesc:   record[2],
			Sku:         record[3],
			Uprice:      Uprice1,
			Availqty:    Availqty1,
			Minstockqty: Minstockqty1,
		}

		resLineDis = append(resLineDis, resLine1)

	}

	if len(resLineDis) == 0 {

		tmpStrNorecs := "No Records in the file"
		sharvar.Tpl.ExecuteTemplate(res, "prodmastresultnorecs.gohtml", tmpStrNorecs)
	} else {

		sharvar.Tpl.ExecuteTemplate(res, "prodmastresult.gohtml", resLineDis)
	}

}

//readProdRecs - read records from the file
func readProdRecs(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	//ignore header record
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

//StockLvlReport - upload product master
func (h1 *ProdHandler) StockLvlReport(res http.ResponseWriter, req *http.Request) {

	var ProdData = map[string]*ProdDet{}

	ProdData, _ = GetProdAllRecsDB()

	if len(ProdData) == 0 {
		tmpStrNorecs := "No Records in the database"
		sharvar.Tpl.ExecuteTemplate(res, "prodstocklvlnorecs.gohtml", tmpStrNorecs)

	} else {

		sharvar.Tpl.ExecuteTemplate(res, "prodstocklvl.gohtml", ProdData)
	}

}

// Database functions
//GetProdRecDB will extract product record from database
func GetProdRecDB(prodIDPar string) (ProdDet, int) {

	db := sharvar.ConnectDB()
	h := NewProdHandler(db)

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	var prodTabRec = ProdDet{}

	var (
		ProdID1      string
		Cat1         string
		Shortdesc1   string
		Sku1         string
		Uprice1      float64
		Pcur1        string
		Availqty1    int
		Minstockqty1 int
	)

	getProdSql := `SELECT prodid,cat,shortdesc,sku, uprice,pcur,availqty,minstockqty  
							FROM prodtab 
							WHERE prodid=?;`

	row := h.db.QueryRow(getProdSql, prodIDPar)
	err := row.Scan(&ProdID1, &Cat1, &Shortdesc1, &Sku1, &Uprice1, &Pcur1, &Availqty1, &Minstockqty1)

	//	fmt.Println("<<< getProdRecDB errResInsertPreExe ", err)

	if err != sql.ErrNoRows {

		prodTabRec = ProdDet{
			ProdID:      ProdID1,
			Cat:         Cat1,
			Shortdesc:   Shortdesc1,
			Sku:         Sku1,
			Uprice:      Uprice1,
			Pcur:        Pcur1,
			Availqty:    Availqty1,
			Minstockqty: Minstockqty1,
		}

		return prodTabRec, 1

	} else {

		return prodTabRec, 0
	}

}

//AddProdRecDB will add new product record to database
func (h1 *ProdHandler) AddProdRecDB(msg ProdDet) error {

	if err := h1.db.Ping(); err != nil {
		panic(err)
	}

	insertPre, errInsertPre := h1.db.Prepare(`INSERT INTO prodtab 
							  (ProdID,cat,shortdesc,sku,uprice,pcur,availqty,minstockqty  )
								VALUES ( ?, ?, ?, ?, ?,?,?,? );`)

	//

	if errInsertPre != nil {
		panic(errInsertPre)
	}

	_, errResInsertPreExe := insertPre.Exec(msg.ProdID, msg.Cat, msg.Shortdesc, msg.Sku,
		msg.Uprice, "SGD", msg.Availqty, msg.Minstockqty)

	//	fmt.Println("<<< AddProdRecDB errResInsertPreExe ", errResInsertPreExe)
	insertPre.Close()
	return errResInsertPreExe
}

//ModProdRecDB will update product record in the database
func (h1 *ProdHandler) ModProdRecDB(msg ProdDet) error {

	if err := h1.db.Ping(); err != nil {
		panic(err)
	}

	updatePre, errUpdatePre := h1.db.Prepare(`update prodtab
											set cat=?,shortdesc=?,sku=?,uprice=?,availqty=?,minstockqty=?
											where prodid=?;`)
	if errUpdatePre != nil {
		panic(errUpdatePre)
	}

	_, errsUpdatePreExe := updatePre.Exec(msg.Cat, msg.Shortdesc, msg.Sku,
		msg.Uprice, msg.Availqty, msg.Minstockqty, msg.ProdID)

	updatePre.Close()

	return errsUpdatePreExe

}

//DelProdRecDB will delete product record from database
func (h1 *ProdHandler) DelProdRecDB(prodIDPar string) error {

	if err := h1.db.Ping(); err != nil {
		panic(err)
	}

	delPre, _ := h1.db.Prepare(`delete from prodtab where ProdID=?;`)

	_, errDelPreExe := delPre.Exec(prodIDPar)

	delPre.Close()

	return errDelPreExe
}

//GetProdAllRecsDB will extract product record from database
func GetProdAllRecsDB() (ProdData, error) {

	db := sharvar.ConnectDB()
	h := NewProdHandler(db)

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	var ProdData1 = map[string]*ProdDet{}

	var (
		ProdID1      string
		Cat1         string
		Shortdesc1   string
		Sku1         string
		Uprice1      float64
		Pcur1        string
		Availqty1    int
		Minstockqty1 int
	)

	getProdSql := `SELECT prodid,cat,shortdesc,sku, uprice,pcur,availqty,minstockqty  
							FROM prodtab ;`

	rows, err := h.db.Query(getProdSql)

	for rows.Next() {
		err2 := rows.Scan(&ProdID1, &Cat1, &Shortdesc1, &Sku1, &Uprice1, &Pcur1, &Availqty1, &Minstockqty1)

		if err2 != nil {

			panic(err2)
		}

		var prodTabRec = ProdDet{}

		prodTabRec = ProdDet{
			ProdID:      ProdID1,
			Cat:         Cat1,
			Shortdesc:   Shortdesc1,
			Sku:         Sku1,
			Uprice:      Uprice1,
			Pcur:        Pcur1,
			Availqty:    Availqty1,
			Minstockqty: Minstockqty1,
		}

		ProdData1[ProdID1] = &prodTabRec

	}
	rows.Close()
	return ProdData1, err
}
