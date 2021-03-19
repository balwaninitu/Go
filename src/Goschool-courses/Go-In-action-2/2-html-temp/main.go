package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("param1")
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, param1)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
