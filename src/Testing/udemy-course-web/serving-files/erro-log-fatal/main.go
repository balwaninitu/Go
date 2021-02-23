package main

import (
	"io"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src = "/Dog.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {

	f, err := os.Open("Dog.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)

}

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/Dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}
