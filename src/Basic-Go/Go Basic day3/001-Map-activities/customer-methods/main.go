package main

import "fmt"

func main() {

	customer1 := customer{
		firstName: "Micheal",
		lastName:  "Jordan",
		username:  "MJ2020",
		password:  "1234567",
		email:     "MJ2020@gmail.com",
		phone:     12345678,
		address:   "18227 Capstan Greens Road Cornelius, NC 28031.",
	}

	customer1.printAllInfo()

	fmt.Println(customer1.printUserCredential())
	fmt.Println(customer1.printAddress())

}
