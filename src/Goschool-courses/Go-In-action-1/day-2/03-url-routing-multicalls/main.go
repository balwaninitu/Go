package main

import (
	"io"
	"net/http"
)

type goMenu int

func (m goMenu) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		io.WriteString(w, "Welcome to Go menu")
	case "/basic":
		io.WriteString(w, "Welcome to Go basic")
	case "/advance":
		io.WriteString(w, "Welcome to Go advance")

	}
}
func main() {
	var a goMenu
	http.ListenAndServe(":8080", a)

}
