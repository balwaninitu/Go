package main

import (
	"html/template"
	"log"
	"net/http"
)

type person struct {
	FirstName string
	LastName  string
	Grade     string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

}

func main() {

	http.HandleFunc("/", form)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func form(w http.ResponseWriter, req *http.Request) {
	firstname := req.FormValue("firstname")
	lastname := req.FormValue("lastname")
	grade := req.FormValue("grade")

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{firstname, lastname, grade})
	if err != nil {
		log.Fatalln(err)
	}

}
