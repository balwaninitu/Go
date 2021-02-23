package main

import (
	"io"
	"net/http"
)

type dog int
type cat int

func (d dog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bow bow bow!!")
}

func (c cat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "meow meow meow!!")
}

func main() {
	var d dog
	var c cat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
