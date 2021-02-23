package main

import (
	"fmt"
	"net/http"
)

type riddham int

func (r riddham) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Ridu's key", "this is from Riddham")
	w.Header().Set("content-Type", "text/html; charset=utf8")
	fmt.Fprintf(w, "<h1>Because I am Happy!!<h1>")

}

func main() {
	var rid riddham

	http.ListenAndServe(":8080", rid)
}
