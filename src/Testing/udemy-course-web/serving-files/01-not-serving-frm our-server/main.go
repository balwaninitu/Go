package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-type", "text/html; chaeset = utf-8")

	io.WriteString(w, `<img src="/dog.jpg">`)
}

func main() {

	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}
