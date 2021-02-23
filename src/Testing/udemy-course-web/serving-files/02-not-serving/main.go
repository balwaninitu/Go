package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="butterfly.jpg"/>`)
}

func main() {

	http.HandleFunc("/", dog)
	http.Handle("/butterfly.jpg", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":8080", nil)
}
