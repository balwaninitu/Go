//Package users is desgined to register new users, update/delete/view existing users details
//It uses github's satori to generate a unique ID to keep track of customers' session
//All passwords are created using native bcruypt library
package users

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"golivejd/sharvar"
	"os"

	//	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//User struc is used to store list of registered customers
type User struct {
	Username     string
	Usertype     string //admin with all previleges or normal with less previleges
	Password     string //string pwd - wont be stored
	HashPassword []byte //hashed pwd

	First   string
	Last    string
	Address string
	Tel     string
	EmailID string

	RecFlag string            // Add, del, dis, modify
	Errors  map[string]string // collect form fields data entry error msgs...
}

//MapUsers is a map used to store registered customers' details including hased passwords
var MapUsers = map[string]User{}

//MapSessions is a map to capture logined users sessions for security reasons
var MapSessions = map[string]string{}

//var MapSessions = map[string]User{}

//IDErrMsg is data structure to collect data entry error msgs for the purpose of data sanity
type IDErrMsg struct {
	Username string
	Password string
	RecFlag  string // Add, del, dis, modify

	Errors map[string]string
}

//
//UsersHandler is a handle to opened database connection
type UsersHandler struct {
	db *sql.DB
}

//NewUsersHandler helps to call user functions using the open db handle
func NewUsersHandler(db *sql.DB) *UsersHandler {
	return &UsersHandler{
		db: db,
	}
}

//func BookFormValidate is used to validate user details during data entry
func (msg *User) UserFormValidate() bool {

	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Username) == "" {
		msg.Errors["Username"] = "Please enter Username"
	}
	// if strings.TrimSpace(msg.Usertype) == "" {
	// 	msg.Errors["Usertype"] = "Please enter Usertype(admin/normal"
	// }
	if strings.TrimSpace(msg.Password) == "" {
		msg.Errors["Password"] = "Please enter Password (will not stored in our DB)"

	} else if !(len(strings.TrimSpace(msg.Password)) > 1) {
		msg.Errors["Password"] = "must be 2 or more characters long"
	}

	sharvar.Mu.Lock()

	//get from database
	_, tRows := GetUserRecDB(msg.Username)
	//	_, ok := MapUsers[msg.Username]
	sharvar.Mu.Unlock()

	if tRows > 0 {
		msg.Errors["Username"] = "Username already exists"
	}

	return len(msg.Errors) == 0
}

//func ValidateID will be used to verify users's login/pwd entries
func (msg *IDErrMsg) ValidateID() bool {

	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Username) == "" {
		msg.Errors["Username"] = "Enter Username"

	}
	if strings.TrimSpace(msg.Password) == "" {
		msg.Errors["Password"] = "Enter Password"

	}

	if strings.TrimSpace(msg.Username) != "" && strings.TrimSpace(msg.Password) != "" {
		// check if user exist with username
		sharvar.Mu.Lock()

		//get from database
		tmpTabUsr, tRows := GetUserRecDB(msg.Username)

		//	_, ok := MapUsers[msg.Username]
		sharvar.Mu.Unlock()

		if tRows == 0 {
			msg.Errors["Username"] = "Entered credential does not exist"
		} else {

			// Matching of password entered

			err := bcrypt.CompareHashAndPassword(tmpTabUsr.HashPassword, []byte(strings.TrimSpace(msg.Password)))
			if err != nil {
				msg.Errors["Username"] = "Entered credential does not exist"
				//			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
				//			return
			}
		}
	}

	return len(msg.Errors) == 0
}

//func checkUsrID will be used to validate booking ID during booking display/updation/deletion
func (msg *IDErrMsg) checkUsrID() bool {

	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Username) == "" {
		msg.Errors["Username"] = "Enter Username"

	}

	if strings.TrimSpace(msg.Username) != "" {

		sharvar.Mu.Lock()
		// check if user exist with username in database

		_, tRows := GetUserRecDB(msg.Username)
		//		_, ok := MapUsers[msg.Username]
		sharvar.Mu.Unlock()

		if tRows == 0 {
			msg.Errors["Username"] = "Username does not exist"
		}
	}

	return len(msg.Errors) == 0
}

//func AddUsr will be used to register coustomers
func (h *UsersHandler) AddUsr(res http.ResponseWriter, req *http.Request) {

	_ = GetUser(res, req)
	if !AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	if req.Method == http.MethodPost {

		//do  bcrypt conversion
		trimwpwd := strings.TrimSpace(req.FormValue("Password"))
		bPassword, _ := bcrypt.GenerateFromPassword([]byte(trimwpwd), bcrypt.MinCost)

		msg := User{
			Username: req.PostFormValue("Username"),
			//	Usertype:     req.PostFormValue("Usertype"),
			Password:     trimwpwd,
			HashPassword: bPassword,
			First:        req.PostFormValue("First"),
			Last:         req.PostFormValue("Last"),
			Tel:          req.PostFormValue("Tel"),
			EmailID:      req.PostFormValue("EmailID"),
			Address:      req.PostFormValue("Address"),

			RecFlag: "add",
		}

		if msg.UserFormValidate() == false {
			sharvar.Tpl.ExecuteTemplate(res, "usradd.gohtml", msg)
			return
		}
		sharvar.Mu.Lock()

		//get rec from database
		_, tRows := GetUserRecDB(msg.Username)
		//_, ok := MapUsers[req.FormValue("Username")]

		sharvar.Mu.Unlock()

		if tRows == 0 {

			sharvar.Mu.Lock()

			//store into database table
			_ = h.AddUserRecDB(msg)

			sharvar.Mu.Unlock()

			tmp := "<br>" + "<b>" + "Username" + " ( " + req.FormValue("Username") + " ) " + "added" + "</b>" + "<br>"
			_, _ = fmt.Fprintln(res, tmp)

		} else {

			tmp := "<br>" + "<b>" + " ( " + req.FormValue("Username") + " ) " + " already exists - can not add" + "</b>" + "<br>"
			_, _ = fmt.Fprintln(res, tmp)

		}

	}

	sharvar.Tpl.ExecuteTemplate(res, "usradd.gohtml", nil)
}

//func DisUsr will be used to display existing users details
func (h *UsersHandler) DisUsr(res http.ResponseWriter, req *http.Request) {

	_ = GetUser(res, req)
	if !AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	var (
		Username string
		tmp      string
	)

	if req.Method == http.MethodPost {

		msg := &IDErrMsg{
			Username: req.FormValue("Username"),
			RecFlag:  "dis",
		}

		if msg.checkUsrID() == false {
			sharvar.Tpl.ExecuteTemplate(res, "getusrWeb.gohtml", msg)
			return
		}
		Username = req.FormValue("Username")

	}

	if Username != "" {

		//get details from database
		UsrTabRec, tRows := GetUserRecDB(Username)

		if tRows > 0 {

			//get user rec from the database and copy the details into existing map user variable

			sharvar.Mu.Lock()
			tmpUsrDet := UsrTabRec

			// MapUsers[Username] = tmpUsrTabRec
			// tmpUsrDet, _ := MapUsers[Username]
			sharvar.Mu.Unlock()

			sharvar.Tpl.ExecuteTemplate(res, "usrdis.gohtml", tmpUsrDet)

		} else {

			tmp = "<br>" + "<b>" + "Username" + " ( " + req.FormValue("Username") + " ) " + " does not exists - can not display" + "</b>" + "<br>"
			_, _ = fmt.Fprintln(res, tmp)
		}
	} else {

		sharvar.Tpl.ExecuteTemplate(res, "getusrWeb.gohtml", nil)

	}

}

//func DelUsr & func DelUsrDetail are meant for deleting a users
func (h *UsersHandler) DelUsr(res http.ResponseWriter, req *http.Request) {

	_ = GetUser(res, req)
	if !AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if err := h.db.Ping(); err != nil {
		panic(err)
	}
	var Username string

	if req.Method == http.MethodPost {

		msg := &IDErrMsg{
			Username: req.FormValue("Username"),
			RecFlag:  "dis",
		}

		if msg.checkUsrID() == false {
			sharvar.Tpl.ExecuteTemplate(res, "getusrWeb.gohtml", msg)
			return
		}

		Username = req.FormValue("Username")
	}

	if Username != "" {

		sharvar.Mu.Lock()
		// get from database
		UsrTabRec, tRows := GetUserRecDB(Username)
		tmpUsrDet := UsrTabRec

		//tmpUsrDet, ok := MapUsers[Username]
		sharvar.Mu.Unlock()

		if tRows > 0 {

			sharvar.Tpl.ExecuteTemplate(res, "usrdel.gohtml", tmpUsrDet)

		} else {

			tmp := "<br>" + "<b>" + "Username" + " ( " + req.FormValue("Username") + " ) " + " does not exist - can not delete" + "</b>" + "<br>"
			_, _ = fmt.Fprintln(res, tmp)
		}
	} else {

		sharvar.Tpl.ExecuteTemplate(res, "getusrWeb.gohtml", nil)
	}

}

//func ModUsr & func ModUsrDetail are meant to modify existing users details
func (h *UsersHandler) ModUsr(res http.ResponseWriter, req *http.Request) {

	_ = GetUser(res, req)
	if !AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	var (
		Username string
	)

	if req.Method == http.MethodPost {

		msg := &IDErrMsg{
			Username: req.FormValue("Username"),
			RecFlag:  "mod",
		}

		if msg.checkUsrID() == false {
			sharvar.Tpl.ExecuteTemplate(res, "getusrWeb.gohtml", msg)
			return
		}

		Username = req.FormValue("Username")

	}

	sharvar.Mu.Lock()
	// get from database for screen field display
	UsrTabRec, tRows := GetUserRecDB(Username)
	tmpUsrDet := UsrTabRec

	//	tmpUsrDet, ok := MapUsers[Username]
	sharvar.Mu.Unlock()

	if tRows > 0 {

		sharvar.Tpl.ExecuteTemplate(res, "usrmod.gohtml", tmpUsrDet)

	} else {

		sharvar.Tpl.ExecuteTemplate(res, "getusrWeb.gohtml", nil)

	}

}

//func DelUsrDetail & func DelUsr are meant for deleting a users
func (h *UsersHandler) DelUsrDetail(res http.ResponseWriter, req *http.Request) {

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	Username := req.FormValue("Username")

	sharvar.Mu.Lock()
	//	delete(MapUsers, Username)
	// delete from database
	_ = h.DelUserRecDB(Username)
	sharvar.Mu.Unlock()

	tmp := "<br>" + "<b>" + "Username" + " ( " + req.FormValue("Username") + " ) " + "deleted" + "</b>" + "<br>"
	_, _ = fmt.Fprintln(res, tmp)

	sharvar.Tpl.ExecuteTemplate(res, "usrdel.gohtml", Username)

	myUser := GetUser(res, req)
	sharvar.Tpl.ExecuteTemplate(res, "index.gohtml", myUser)

}

//func ModUsrDetail & func ModUsr are meant to modify existing users details
func (h *UsersHandler) ModUsrDetail(res http.ResponseWriter, req *http.Request) {

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	if req.Method == http.MethodPost {

		usr1 := User{
			Username: req.PostFormValue("Username"),
			//	Usertype: req.PostFormValue("Usertype"),
			Password: req.PostFormValue("Password"),
			//			HashPassword: bPassword,
			First: req.PostFormValue("First"),
			Last:  req.PostFormValue("Last"),

			Tel:     req.PostFormValue("Tel"),
			EmailID: req.PostFormValue("EmailID"),
			Address: req.PostFormValue("Address"),

			RecFlag: "mod",
		}

		sharvar.Mu.Lock()
		usr1.Password = ""
		//MapUsers[req.FormValue("Username")] = usr1
		//save to database (modified fields)
		_ = h.ModUserRecDB(usr1)

		sharvar.Mu.Unlock()

		tmp := "<br>" + "<b>" + "Username" + " ( " + req.PostFormValue("Username") + " ) " + "modified" + "</b>" + "<br>"
		_, _ = fmt.Fprintln(res, tmp)
		//	}
	}
	myUser := GetUser(res, req)
	sharvar.Tpl.ExecuteTemplate(res, "index.gohtml", myUser)

}

//func Signup is used to register new users
//this will also create sessions and cookies
func (h *UsersHandler) Signup(res http.ResponseWriter, req *http.Request) {

	if AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var myUser User
	// process form submission
	if req.Method == http.MethodPost {

		trimwpwd := strings.TrimSpace(req.FormValue("Password"))
		bPassword, err := bcrypt.GenerateFromPassword([]byte(trimwpwd), bcrypt.MinCost)

		//		bPassword, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
			return
		}

		msg := User{
			Username: req.PostFormValue("Username"),
			//		Usertype:     req.PostFormValue("Usertype"),
			Password:     trimwpwd,
			HashPassword: bPassword,
			First:        req.PostFormValue("First"),
			Last:         req.PostFormValue("Last"),

			Tel:     req.PostFormValue("Tel"),
			EmailID: req.PostFormValue("EmailID"),
			Address: req.PostFormValue("Address"),

			RecFlag: "add",
		}

		if msg.UserFormValidate() == false {
			sharvar.Tpl.ExecuteTemplate(res, "signup.gohtml", msg)
			return
		}

		sharvar.Mu.Lock()
		//get rec from database
		_, tRows := GetUserRecDB(msg.Username)

		//_, ok := MapUsers[req.FormValue("Username")]
		sharvar.Mu.Unlock()

		if tRows == 0 { // new users

			sharvar.Mu.Lock()

			MapUsers[req.FormValue("Username")] = msg

			//store into database table
			_ = h.AddUserRecDB(msg)

			sharvar.Mu.Unlock()

			// create session
			id, _ := uuid.NewV4()
			myCookie := &http.Cookie{
				Name: "myCookie",
				//		Expires:  expireCookie,
				HttpOnly: true,
				Path:     "/",
				//				Domain: "127.0.0.1",
				Secure: true,

				Value: id.String(),
			}
			http.SetCookie(res, myCookie)
			MapSessions[myCookie.Value] = msg.Username

			// redirect to main index
			http.Redirect(res, req, "/", http.StatusSeeOther)

			// tmp := "<br>" + "<b>" + "Username" + " ( " + req.FormValue("Username") + " ) " + "added" + "</b>" + "<br>"
			// _, _ = fmt.Fprintln(res, tmp)
			return

		} else {

			// tmp := "<br>" + "<b>" + " ( " + req.FormValue("Username") + " ) " + " already exists - can not add" + "</b>" + "<br>"
			// _, _ = fmt.Fprintln(res, tmp)

		}
	}
	sharvar.Tpl.ExecuteTemplate(res, "signup.gohtml", myUser)

}

//func Login allows a users to log into the application
//this will also setup related sessions and cookies
func (h *UsersHandler) Login(res http.ResponseWriter, req *http.Request) {

	if AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		msg := IDErrMsg{
			Username: req.FormValue("Username"),
			Password: req.FormValue("Password"),
		}

		if msg.ValidateID() == false {
			sharvar.Tpl.ExecuteTemplate(res, "login.gohtml", msg)
			return
		}
		//read from database table
		tmpTabUsr, tRows := GetUserRecDB(req.FormValue("Username"))

		var myUser User //

		if tRows > 0 {

			myUser = tmpTabUsr

		}

		myUser1 := User{
			Username: tmpTabUsr.Username,
			//		Usertype: tmpTabUsr.Usertype,
			First: tmpTabUsr.First,
			Last:  tmpTabUsr.Last,

			RecFlag: "add",
		}

		MapUsers[req.FormValue("Username")] = myUser1
		// create session
		id, _ := uuid.NewV4()
		myCookie := &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}

		http.SetCookie(res, myCookie)
		MapSessions[myCookie.Value] = myUser.Username

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return

	}

	sharvar.Tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

//func Logout allows a users to logout of the application
//this will also delete related sessions and cookies
func (h *UsersHandler) Logout(res http.ResponseWriter, req *http.Request) {

	if !AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	myCookie, _ := req.Cookie("myCookie")
	// delete the session
	delete(MapSessions, myCookie.Value)
	// remove the cookie
	myCookie = &http.Cookie{
		Name:   "myCookie",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, myCookie)

	http.Redirect(res, req, "/", http.StatusSeeOther)
}

//func GetUser to keep track of session
func GetUser(res http.ResponseWriter, req *http.Request) User {

	// get current session cookie
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		id, _ := uuid.NewV4()
		myCookie = &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}

	}
	http.SetCookie(res, myCookie)

	// if the user exists already, get user
	var myUser User
	if Username, ok := MapSessions[myCookie.Value]; ok {
		myUser = MapUsers[Username]
	}

	return myUser
}

// func AlreadyLoggedIn is a middleware to prevent multiple logins of customers
func AlreadyLoggedIn(req *http.Request) bool {

	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		return false

	}

	Username := MapSessions[myCookie.Value]

	_, ok := MapUsers[Username]

	return ok
}

//UplUsrMast - upload user master
func (h *UsersHandler) UplUsrMast(res http.ResponseWriter, req *http.Request) {

	_ = GetUser(res, req)
	if !AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	file := "C:\\Projects\\Go\\src\\golivejd\\assets\\usermast.csv"
	var (
		msg       = User{}
		UpdStatus string
	)

	type resLine struct {
		UpdStatus string
		Username  string
		First     string
		Last      string
		Address   string
		Tel       string
		EmailID   string
	}

	var resLineDis = []resLine{}

	records, err := readUserRecs(file)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {

		//do  bcrypt conversion
		trimwpwd := strings.TrimSpace(record[2])
		bPassword, _ := bcrypt.GenerateFromPassword([]byte(trimwpwd), bcrypt.MinCost)

		msg = User{
			Username:     strings.TrimSpace(record[0]),
			HashPassword: bPassword,
			First:        record[2],
			Last:         record[3],
			Address:      record[4],
			Tel:          record[5],
			EmailID:      record[6],
		}

		tmpUsrDet, tRows := GetUserRecDB(msg.Username)

		if tRows == 0 {

			h.AddUserRecDB(msg)
			UpdStatus = "New Rec Uploaded"

		} else {

			if tmpUsrDet.Username == msg.Username && tmpUsrDet.First == msg.First && tmpUsrDet.Last == msg.Last && tmpUsrDet.Tel == msg.Tel && tmpUsrDet.Address == msg.Address && tmpUsrDet.EmailID == msg.EmailID {
				UpdStatus = "Rec No Changes Made"

			} else {

				_ = h.ModUserRecDB(msg)
				UpdStatus = "Existing Rec Updated"

			}

		}

		resLine1 := resLine{
			UpdStatus: UpdStatus,
			Username:  strings.TrimSpace(record[0]),
			First:     record[2],
			Last:      record[3],
			Address:   record[4],
			Tel:       record[5],
			EmailID:   record[6],
		}

		resLineDis = append(resLineDis, resLine1)

	}

	if len(resLineDis) == 0 {

		tmpStrNorecs := "No Records in the file"
		sharvar.Tpl.ExecuteTemplate(res, "usrmastresultnorecs.gohtml", tmpStrNorecs)
	} else {

		sharvar.Tpl.ExecuteTemplate(res, "usrmastresult.gohtml", resLineDis)

	}

	// if tmpUsrDet.Username != msg.Username {

	// 	} else if tmpUsrDet.First != msg.First {

	// 	} else if tmpUsrDet.Last != msg.Last {
	// 	} else if tmpUsrDet.Tel != msg.Tel {
	// 	} else if tmpUsrDet.EmailID != msg.EmailID {
	// 	}

}

//readUserRecs - read records from the file
func readUserRecs(fileName string) ([][]string, error) {

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

//
//GetUserRecDB will extract user record from database
func GetUserRecDB(unamePar string) (User, int) {

	db := sharvar.ConnectDB()
	h := NewUsersHandler(db)

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	var UserTabRec = User{}

	var (
		username1 string
		//	usertype1 string
		hpwd1    []byte
		first1   string
		last1    string
		Address1 string
		Tel1     string
		EmailID1 string
	)

	getUserSql := `SELECT Username, HashPassword, First,Last,tel,emailid,address FROM userstab WHERE username=?;`

	row := h.db.QueryRow(getUserSql, unamePar)
	err := row.Scan(&username1, &hpwd1, &first1, &last1, &Tel1, &EmailID1, &Address1)

	if err != sql.ErrNoRows {

		UserTabRec = User{
			Username: username1,
			//		Usertype:     usertype1,
			HashPassword: hpwd1,
			First:        first1,
			Last:         last1,
			Address:      Address1,
			Tel:          Tel1,
			EmailID:      EmailID1,
		}

		return UserTabRec, 1

	} else {

		return UserTabRec, 0
	}

}

//AddUserRecDB will add new user record to database
func (h *UsersHandler) AddUserRecDB(msg User) error {

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	insertPre, errInsertPre := h.db.Prepare(`INSERT INTO userstab (username,hashPassword,first,last,tel,emailid,address) 
											VALUES (  ?, ?, ?, ?, ?,?,? );`)

	if errInsertPre != nil {
		panic(errInsertPre)
	}
	//	resInsertPre, errResInsertPreExe := insertPre.Exec(msg.Username, msg.Usertype, msg.HashPassword, msg.First, msg.Last)
	_, errResInsertPreExe := insertPre.Exec(msg.Username, msg.HashPassword, msg.First, msg.Last, msg.Tel, msg.EmailID, msg.Address)

	insertPre.Close()
	return errResInsertPreExe
}

//ModUserRecDB will update new user record in the database
func (h *UsersHandler) ModUserRecDB(msg User) error {

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	// updatePre, errUpdatePre := h.db.Prepare(`update userstab
	// 							set Usertype=?, First=?, Last=?
	// 							where Username=?;`)

	updatePre, errUpdatePre := h.db.Prepare(`update userstab
								set First=?, Last=?, tel=?,emailid=?,address=?
								where Username=?;`)

	if errUpdatePre != nil {
		panic(errUpdatePre)
	}

	//	resUpdatePre, errsUpdatePreExe := updatePre.Exec(msg.Usertype, msg.First, msg.Last, msg.Username)
	_, errsUpdatePreExe := updatePre.Exec(msg.First, msg.Last, msg.Tel, msg.EmailID, msg.Address, msg.Username)

	updatePre.Close()

	return errsUpdatePreExe

}

//DelUserRecDB will delete user record from database
func (h *UsersHandler) DelUserRecDB(Username string) error {

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

	// delPre, errDelPre := h.db.Prepare(`delete from userstab
	// where Username=?;`)

	delPre, errDelPre := h.db.Prepare(`delete from userstab
						where Username=?;`)

	if errDelPre != nil {
		panic(errDelPre)
	}

	//resDelPre, errDelPre := delPre.Exec(Username)
	_, errDelPreExe := delPre.Exec(Username)

	delPre.Close()

	return errDelPreExe

} //UserRecExistsDB is a helper function to make to avoid duplicate records in the database
func userRecExistsDB(Username string) {

	db := sharvar.ConnectDB()
	h := NewUsersHandler(db)

	if err := h.db.Ping(); err != nil {
		panic(err)
	}

}
