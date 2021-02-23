package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type riddham int

func (r riddham) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	var rid riddham

	http.ListenAndServe(":8080", rid)

}
