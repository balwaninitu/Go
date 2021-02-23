package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
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
		Submissions map[string][]string
		Header      http.Header
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
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
