package main

import (
	"fmt"
)

func main() {

	const (
		arithFunc1 = "+"
		arithFunc2 = "-"
		arithFunc3 = "*"
		arithFunc4 = "/"
	)

	var inputVal1, inputVal2 int

	var arithFunc string

	fmt.Println("Input a value from 0, 1, 2, 3 …. 9")

	fmt.Scanln(&inputVal1)

	fmt.Println("Select one of arithmatic function among +, -, * or /")

	fmt.Scanln(&arithFunc)

	fmt.Println("Input a value from 0, 1, 2, 3 …. 9")

	fmt.Scanln(&inputVal2)

	if arithFunc == arithFunc1 {

		fmt.Println("Final result:", inputVal1+inputVal2)

	} else if arithFunc == arithFunc2 {

		fmt.Println("Final result:", inputVal1-inputVal2)

	} else if arithFunc == arithFunc3 {

		fmt.Println("Final result:", inputVal1*inputVal2)
	} else if arithFunc == arithFunc4 {

		fmt.Println("Final result:", inputVal1/inputVal2)
	}

}
