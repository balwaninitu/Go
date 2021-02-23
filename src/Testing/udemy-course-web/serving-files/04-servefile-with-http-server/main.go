package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("ContentType", `"text/html; charser=utf-8`)
	io.WriteString(w, `<img src = "/Dog.jpg ">`)
}

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}
