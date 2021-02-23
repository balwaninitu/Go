package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", uRLcall)
	http.Handle("/faviocon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func uRLcall(w http.ResponseWriter, req *http.Request) {
	passValue := req.FormValue("query")
	fmt.Fprintf(w, "Value is:"+passValue)
}
