//Admin functions
package admin

import (
	"net/http"
	"projectGoLive/application/config"
	"projectGoLive/application/user_db"
)

type adminstruct struct {
	Operation   string
	Mainmessage []string
	UserInfo    []user_db.UserDetails
}

//---------------------------------------------------------------------------
// Functions for admin
//---------------------------------------------------------------------------
// This method is used to perform all admin functions
// func AdminHandler(w http.ResponseWriter, req *http.Request) {
// 	if !server.ActiveSession(w, req) {
// 	http.Redirect(w, req, "/admin", http.StatusSeeOther)
// 	return
// 	}

// 	config.TPL.ExecuteTemplate(w, "admin.gohtml", nil)
// }s

func AdminHandler(w http.ResponseWriter, req *http.Request) {
	config.TPL.ExecuteTemplate(w, "admin.gohtml", nil)
}

func ViewSller(w http.ResponseWriter, req *http.Request) {
	sellerdetails, _ := user_db.GetRecords(config.DB)

	sellerMessage := adminstruct{
		Operation:   "view",
		Mainmessage: nil,
		UserInfo:    sellerdetails,
	}
	sellerMessage.Mainmessage = append(sellerMessage.Mainmessage, "Here are Seller profile details :")

	config.TPL.ExecuteTemplate(w, "admin.gohtml", sellerMessage)
}
