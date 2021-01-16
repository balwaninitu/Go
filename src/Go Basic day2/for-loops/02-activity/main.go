package main

import "fmt"

func main() {

	var inputNum int

	fmt.Println("Please enter number")

	fmt.Scanln(&inputNum)

	if inputNum%2 == 0 {

		fmt.Println("Number is even")
	} else {

		fmt.Println("Number is odd")
	}

}
