package main

import (
	"io"
	"net/http"
)

type riddham int

func (r riddham) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "bhow Bhow Bhow!!!")
	case "/cat":
		io.WriteString(w, "Meow meow meow!!!")
	}
}

func main() {
	var rid riddham
	http.ListenAndServe(":8080", rid)
}
