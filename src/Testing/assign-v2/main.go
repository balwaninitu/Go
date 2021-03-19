package main

import (
	"net/http"
	"text/template"
)

type doctorDetails struct {
	Id   string
	Name string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/name", name)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	doctor := getDoctor(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", doctor)

}

func getDoctor(res http.ResponseWriter, req *http.Request) doctorDetails {
	data := doctorDetails{
		Name: "name",
	}

	return data
}

func name(res http.ResponseWriter, req *http.Request) {
	// process form submission
	if req.Method == http.MethodPost {
		name := req.FormValue("name")
		tpl.ExecuteTemplate(res, "name.gohtml", name)
	}
}
