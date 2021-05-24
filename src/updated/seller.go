//Seller functions, talks to apiclient.go, server.go and templates
package seller

import (
	"fmt"
	"net/http"
	"projectGoLive/application/apiclient"
	"projectGoLive/application/config"
	"projectGoLive/application/server"
)

type sellerStruct struct {
	Sellername  string
	Operation   string
	Mainmessage []string
	Selleritems []apiclient.ItemsDetails
}

//---------------------------------------------------------------------------
// Functions to display all items added by seller
//---------------------------------------------------------------------------
// This method is used to view the items added by seller
func SellerHandler(w http.ResponseWriter, req *http.Request) {
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := server.GetUser(w, req)

	sellerMessage := sellerStruct{
		Sellername:  user.Username,
		Operation:   "view",
		Mainmessage: nil,
		Selleritems: nil,
	}
	sellerMessage.Mainmessage = append(sellerMessage.Mainmessage, "List of items added: ")

	if user.IsBuyer { // Not possible
		config.Trace.Printf("Incorrect login information! username: %s  is not a seller.", user.Username)
		config.Error.Printf("Incorrect login information! username: %s  is not a seller.", user.Username)
		return
	}

	si, ok := apiclient.GetItem("", user.Username, user.IsBuyer) // get all items for this seller only
	if !ok {
		config.Trace.Printf("Unable to get item data for seller %s \n", user.Username)
		config.Error.Printf("Unable to get item data for seller %s \n", user.Username)
		return
	}

	sellerMessage.Selleritems = si

	config.TPL.ExecuteTemplate(w, "sellertemplate.gohtml", sellerMessage)
}

//---------------------------------------------------------------------------
// Functions to add an item for seller
//---------------------------------------------------------------------------
func AddItemHandler(w http.ResponseWriter, req *http.Request) {
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := server.GetUser(w, req)

	sellerMessage := sellerStruct{
		Sellername:  user.Username,
		Operation:   "add",
		Mainmessage: nil,
		Selleritems: nil,
	}

	// process form submission , when buyer clicks submit
	if req.Method == http.MethodPost {
		fruitname := req.FormValue("fruit")
		quantity := req.FormValue("quantity")
		cost := req.FormValue("cost")
		fmt.Println(fruitname, quantity, cost)

		var item apiclient.ItemsDetails

		if fruitname != "" {
			item.Item = fruitname
			item.Quantity = config.ConvertToInt(quantity)
			item.Cost = config.ConvertToFloat(cost)
			item.Username = user.Username

			ok := apiclient.AddItem(item.Item, user.Username, user.IsBuyer, item)
			if !ok {
				sellerMessage.Mainmessage = append(sellerMessage.Mainmessage, "Unable to add item\nIf item already exists, please update item, else try again!")
			} else {
				http.Redirect(w, req, "/seller", http.StatusSeeOther)
				return
			}

		}
	}
	config.TPL.ExecuteTemplate(w, "sellertemplate.gohtml", sellerMessage)
}

//---------------------------------------------------------------------------
// Functions to update an item added by seller
//---------------------------------------------------------------------------
func UpdateItemHandler(w http.ResponseWriter, req *http.Request) {
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := server.GetUser(w, req)

	sellerMessage := sellerStruct{
		Sellername:  user.Username,
		Operation:   "update",
		Mainmessage: nil,
		Selleritems: nil,
	}
	config.TPL.ExecuteTemplate(w, "sellertemplate.gohtml", sellerMessage)
}

//---------------------------------------------------------------------------
// Functions to delete an item added by seller
//---------------------------------------------------------------------------
func DeleteItemHandler(w http.ResponseWriter, req *http.Request) {
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := server.GetUser(w, req)

	sellerMessage := sellerStruct{
		Sellername:  user.Username,
		Operation:   "delete",
		Mainmessage: nil,
		Selleritems: nil,
	}
	config.TPL.ExecuteTemplate(w, "sellertemplate.gohtml", sellerMessage)
}

//---------------------------------------------------------------------------
// Functions to display profile of seller
//---------------------------------------------------------------------------
func ShowProfile(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := server.GetUser(w, req)

	sellerMessage := sellerStruct{
		Sellername:  user.Username,
		Operation:   "add",
		Mainmessage: nil,
		Selleritems: nil,
	}
	config.TPL.ExecuteTemplate(w, "sellertemplate.gohtml", sellerMessage)
}
