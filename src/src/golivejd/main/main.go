//Package main is the entry point into this application and drives its entire data flow and business logic
//makes use of a couple of packages
package main

import (
	"fmt"
	"golivejd/product"
	"golivejd/sharvar"
	"golivejd/users"
	"net/http"
)

//func init lodes a few sample data for the purpose quick tesing
func init() {

}

func main() {

	//Handle to open db connection
	db := sharvar.ConnectDB()
	//Handle for users package funcions
	h := users.NewUsersHandler(db)

	//Handle for Product package funcions
	h1 := product.NewProdHandler(db)

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", h.Signup)

	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/logout", h.Logout)

	//Product Mgt
	http.HandleFunc("/NewProduct", h1.AddProd)
	http.HandleFunc("/DisProduct", h1.DisProd)
	http.HandleFunc("/ModProduct", h1.ModProd)
	http.HandleFunc("/ModProdRec", h1.ModProdRec)

	http.HandleFunc("/DelProduct", h1.DelProd)
	http.HandleFunc("/DelProdRec", h1.DelProdRec)

	http.HandleFunc("/UplProdMast", h1.UplProdMast)
	http.HandleFunc("/StockLvlReport", h1.StockLvlReport)

	//User Mgt
	http.HandleFunc("/NewUsr", h.AddUsr)
	http.HandleFunc("/DisUsr", h.DisUsr)
	http.HandleFunc("/ModUsr", h.ModUsr)
	http.HandleFunc("/DelUsr", h.DelUsr)

	http.HandleFunc("/ModUsrDet", h.ModUsrDetail)
	http.HandleFunc("/DelUsrDet", h.DelUsrDetail)

	//UplUsrMast

	http.HandleFunc("/UplUsrMast", h.UplUsrMast)

	//Https secure server
	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Println("listening on 8080...")
	//http.ListenAndServeTLS(":8081", "C:\\Projects\\Go\\src\\certs\\cert.pem", "C:\\Projects\\Go\\src\\certs\\key.pem", nil)
	http.ListenAndServe(":8080", nil)
}

//func index is the starting page of this whole online web application
func index(res http.ResponseWriter, req *http.Request) {
	myUser := users.GetUser(res, req)

	sharvar.Tpl.ExecuteTemplate(res, "index.gohtml", myUser)

}
