package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHTML)
	//mux.HandleFunc("/signup", signupHTML)
	//mux.HandleFunc("/login", loginHTML)
	//mux.HandleFunc("/logout", logoutHTML)
	http.ListenAndServe(":5221", mux)
}

func indexHTML(res http.ResponseWriter, req *http.Request) {
	//userx := getUser(res, req)
	//if userx.Username != "" {
	//http.Redirect(res, req, "/explore", http.StatusSeeOther)
	//Warning.Println("Already logged in.")
	//return

	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
