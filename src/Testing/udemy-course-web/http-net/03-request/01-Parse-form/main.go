package main

import (
	"log"
	"net/http"
	"text/template"
)

type riddham int

func (ri riddham) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)

}

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	var rid riddham
	http.ListenAndServe(":8080", rid)

}
