package main

import (
	"io"
	"net/http"
)

func goBasic(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to Go Basic")
}

func goAdvance(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to Go Advance")
}

func goMenu(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to Go Menu")
}

func main() {

	http.Handle("/", http.HandlerFunc(goMenu))
	http.Handle("/advanced", http.HandlerFunc(goAdvance))
	http.Handle("/basic", http.HandlerFunc(goBasic))
	http.ListenAndServe(":8080", nil)
}
