package main

import (
	"io"
	"net/http"
)

//strip prefix only takes file which we want it to take & jot main.go /code file like earlier
func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src = "/resources/Dog.jpg">`)
}

func main() {

	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)

}
