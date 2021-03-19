package main

import (
	"io"
	"net/http"
)

func feature1(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "welcome to feature 1")
}

func feature2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "welcome to feature 2")
}

func main() {
	http.Handle("/javicon.ico", http.NotFoundHandler())
	http.HandleFunc("/feature1", feature1)
	http.HandleFunc("/feature2", feature2)
	http.ListenAndServe(":8080", nil)

}
