package main

import "fmt"

const (
	authenticatedUser1 = "Admin"
	authenticatedUser2 = "Robin"
	authenticatedUser3 = "John"
)

func main() {

	var name string

	fmt.Println("Please enter your name")

	fmt.Scanln(&name)

	if name == authenticatedUser1 {

		fmt.Println("Welcome Admin")
	} else if name == authenticatedUser2 || name == authenticatedUser3 {

		fmt.Println("Welcome")
	} else {

		fmt.Println("Merry Men")
	}

}
