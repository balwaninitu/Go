package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mylogger()
	Logger.Print("This is my test log message")
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("name")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
	 <input type="text" name="name">
	 <input type="submit">
	</form>
	<br>`+v)
}

var (
	Logger *log.Logger
)

func mylogger() {
	var errorLogpath = "error.log"
	var errorFile, err = os.OpenFile(errorLogpath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error", err)
	}
	Logger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
