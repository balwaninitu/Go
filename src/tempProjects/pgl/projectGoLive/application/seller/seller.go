//Seller functions, talks to apiclient.go, server.go and templates
package seller

import (
	"fmt"
	"net/http"
	"pgl/projectGoLive/application/apiclient"
	"pgl/projectGoLive/application/config"
	"pgl/projectGoLive/application/server"
)

type sellerStruct struct {
	Sellername  string
	Location    string
	Operation   string
	Mainmessage []string
	Selleritems []apiclient.ItemsDetails
	Sellerfunc  func(int) int
}

/*
func (s sellerStruct) Inc(x int) int {
	return x + 1
}*/

func increment(x int) int {
	return x + 1
}

/*
func decrement(x int) int {
	return x - 1
}
func multiply(x int, y float64) float64 {
	return float64(x) * y
}

var Fm = template.FuncMap{
	"inc": increment,
	"dec": decrement,
	"mul": multiply,
}
*/

//---------------------------------------------------------------------------
// Functions to display all items added by seller
//---------------------------------------------------------------------------
// This method is used to view the items added by seller
func SellerHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := server.GetUser(w, req)

	sellerMessage := sellerStruct{
		Sellername:  user.Username,
		Location:    user.Location,
		Operation:   "view",
		Mainmessage: nil,
		Selleritems: nil,
		Sellerfunc: func(x int) int {
			return x + 1
		},
	}
	//sellerMessage.Mainmessage = append(sellerMessage.Mainmessage, "Welcome Seller! ")
	sellerMessage.Mainmessage = append(sellerMessage.Mainmessage, "List of items added: ")
	//sellerMessage.Sellername = user.Username
	//sellerMessage.Location = user.Location

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
	//sellerMessage.Sellerfunc = increment
	//sellerMessage.Sellerfunc
	//config.TPL = config.TPL.Funcs(Fm)
	//config.TPL = config.TPL.Funcs(template.FuncMap{"addfunc": increment})
	config.TPL.ExecuteTemplate(w, "sellertemplate.gohtml", sellerMessage)
}

//---------------------------------------------------------------------------
// Functions to add an item for seller
//---------------------------------------------------------------------------
func AddItemHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := server.GetUser(w, req)

	sellerMessage := sellerStruct{
		Sellername:  user.Username,
		Location:    user.Location,
		Operation:   "add",
		Mainmessage: nil,
		Selleritems: nil,
		Sellerfunc: func(x int) int {
			return x + 1
		},
	}

	// process form submission , when buyer clicks submit
	if req.Method == http.MethodPost {
		fruitname := req.FormValue("fruit")
		quantity := req.FormValue("quantity")
		cost := req.FormValue("cost")
		fmt.Println(fruitname, quantity, cost)

		if fruitname != "" {

		}
	}
	config.TPL.ExecuteTemplate(w, "sellertemplate.gohtml", sellerMessage)
}

//---------------------------------------------------------------------------
// Functions to update an item added by seller
//---------------------------------------------------------------------------
func UpdateItemHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := server.GetUser(w, req)

	sellerMessage := sellerStruct{
		Sellername:  user.Username,
		Location:    user.Location,
		Operation:   "update",
		Mainmessage: nil,
		Selleritems: nil,
		Sellerfunc: func(x int) int {
			return x + 1
		},
	}
	config.TPL.ExecuteTemplate(w, "sellertemplate.gohtml", sellerMessage)
}

//---------------------------------------------------------------------------
// Functions to delete an item added by seller
//---------------------------------------------------------------------------
func DeleteItemHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if !server.ActiveSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	user := server.GetUser(w, req)

	sellerMessage := sellerStruct{
		Sellername:  user.Username,
		Location:    user.Location,
		Operation:   "delete",
		Mainmessage: nil,
		Selleritems: nil,
		Sellerfunc: func(x int) int {
			return x + 1
		},
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
		Location:    user.Location,
		Operation:   "add",
		Mainmessage: nil,
		Selleritems: nil,
		Sellerfunc: func(x int) int {
			return x + 1
		},
	}
	config.TPL.ExecuteTemplate(w, "sellertemplate.gohtml", sellerMessage)
}
