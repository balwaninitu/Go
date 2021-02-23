package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src = "Dog.jpg">      `)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "Dog.jpg")
}

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/Dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)

}
