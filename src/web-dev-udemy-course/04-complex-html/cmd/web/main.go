package main

import (
	"fmt"
	"net/http"
	"web-dev-udemy-course/04-complex-html/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting Application on %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
