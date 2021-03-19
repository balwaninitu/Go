package main

import "fmt"

func main() {

	var num int

	fmt.Println("Please enter an integer number")

	fmt.Scanln(&num)

	if num%5 == 0 && num%6 == 0 {

		fmt.Println("Number is divisible by both 5 & 6")
	} else {
		fmt.Println("Number is not divisible by both 5 & 6")

	}

}
