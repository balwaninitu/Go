//Seller functions, talks to apiclient.go, server.go and templates
package seller

import (
	"net/http"
	"projectGoLive/application/config"
	"projectGoLive/application/server"
)

//---------------------------------------------------------------------------
// Functions to display items added by seller
//---------------------------------------------------------------------------
// This method is used to view the items added by seller
func SellerHandler(w http.ResponseWriter, req *http.Request) {
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := server.GetUser(w, req)
	//fmt.Println(user)
	message := []string{}

	message = append(message, "Hello! ")
	message = append(message, user.Username)
	if user.IsBuyer {
		message = append(message, "  You are a Buyer ")
	} else {
		message = append(message, "  You are a Seller ")
	}

	//if !user.IsBuyer {
	//	apiclient.GetItem("", "", true) // get all items

	/*
		if err != nil {
			Trace.Printf("Unable to get appointment data : %v\n", err)
			Error.Printf("Unable to get appointment data : %v\n", err)

			return
		}*/
	//message = msg
	//} else {
	//message =
	//}
	config.TPL.ExecuteTemplate(w, "sellertemplate.gohtml", message)
}
