package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", form)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func form(w http.ResponseWriter, req *http.Request) {
	value := req.FormValue("text")
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w,
		`<form method="POST">
		<input type="text" name="text">
		<input type="First Name">
		<input type="Last Name">
		<input type="Grade">
		</form>
		<br>`+value)
}
