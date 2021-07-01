package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(4, 5)
	fmt.Fprintf(w, "This is sum given from func Page %d\n", sum)
}
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
		return
	}
	fmt.Fprintf(w, "Here is divide result %f ", result)

}

func divideValues(x, y float64) (float64, error) {
	result := x / y

	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	return result, nil
}

func main() {

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Printf("Listening on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
