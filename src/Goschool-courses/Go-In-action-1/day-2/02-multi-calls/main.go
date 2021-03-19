package main

import (
	"fmt"
	"io"
	"net/http"
)

type goBasic int
type goAdvance int
type goMenu int

func (b goBasic) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to Go Basic")
	fmt.Println("menu")
}

func (a goAdvance) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to Go Advance")
}

func (i goMenu) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to Go Menu")
}

func main() {
	var p goBasic
	var q goAdvance
	var r goMenu

	http.Handle("/", r)
	http.Handle("/advanced", q)
	http.Handle("/basic", p)
	http.ListenAndServe(":8080", nil)
}
