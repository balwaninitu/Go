package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bow bow bow!!")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "meow meow meow!!")
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)

}
