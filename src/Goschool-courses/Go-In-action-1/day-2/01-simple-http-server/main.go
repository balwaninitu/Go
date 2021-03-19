package main

import (
	"fmt"
	"net/http"
)

type serveHTTPPage int

func (s serveHTTPPage) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("GoSchoolKey", "This is from GoSchool")
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintf(w, "<h1>Your code is in this func</h1>")
}

func main() {

	var d serveHTTPPage
	http.ListenAndServe(":8080", d)

}
