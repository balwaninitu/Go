package main

import (
	"io"
	"net/http"
)

type goMenu int
type goBasic int
type goAdvance int

func (m goMenu) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to Go Menu")
}

func (b goBasic) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to Go Basic")
}

func (a goAdvance) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to Go Advance")
}

func main() {

	var p goMenu
	var q goBasic
	var r goAdvance

	mux := http.NewServeMux()
	mux.Handle("/", p)
	mux.Handle("/basic", q)
	mux.Handle("/advance", r)
	http.ListenAndServe(":8080", mux)

}
