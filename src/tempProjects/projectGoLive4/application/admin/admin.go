//Admin functions
package admin

import (
	"database/sql"
	"net/http"
	"projectGoLive/application/config"
	"projectGoLive/application/user_db"

	_ "github.com/go-sql-driver/mysql"
)

type ItemsDetails struct {
	Item     string  `json:"Item"`
	Quantity int     `json:"Quantity"`
	Cost     float64 `json:"Cost"`
	Username string  `json:"Username"`
}
type UserDetails struct {
	Username string
	Password string
	Fullname string
	Isbuyer  bool
	Phone    string
	Address  string
	Email    string
}

type adminStruct struct {
	Operation   string
	Mainmessage []string
	Sellername  user_db.UserDetails
}

var db *sql.DB

//---------------------------------------------------------------------------
// Functions for admin
//---------------------------------------------------------------------------
// This method is used to perform all admin functions
func AdminHandler(w http.ResponseWriter, req *http.Request) {
	// if !server.ActiveSession(w, req) {
	// 	http.Redirect(w, req, "/login", http.StatusSeeOther)
	// 	return
	// }

	config.TPL.ExecuteTemplate(w, "admin.gohtml", nil)
}

func ViewSeller(w http.ResponseWriter, r *http.Request) {
	//var sellerItems ItemsDetails

	sellerItems, ok := user_db.GetRecords(db)
	if ok {
		config.TPL.ExecuteTemplate(w, "admin.gohtml", sellerItems)
	}

}

// func UserIndex(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
// 		return
// 	}

// 	rows, err := config.DB.Query("SELECT * FROM UserDetails")
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()
// 	items := make([]UserDetails, 0)
// 	for rows.Next() {
// 		itm := UserDetails{}
// 		err := rows.Scan(&itm.Item, &itm.Quantity, &itm.Cost, &itm.Username)
// 		if err != nil {
// 			http.Error(w, http.StatusText(500), 500)
// 			return
// 		}
// 		items = append(items, itm)
// 	}
// 	if err = rows.Err(); err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	config.TPL.ExecuteTemplate(w, "item.gohtml", items)

// }
