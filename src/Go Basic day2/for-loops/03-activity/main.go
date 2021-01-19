package main

import (
	"fmt"
)

func main() {

	var inputNum1, inputNum2 int

	fmt.Println("Please enter any one number")

	fmt.Scanln(&inputNum1)

	fmt.Println("Please enter one more number")

	fmt.Scanln(&inputNum2)

	if inputNum2 < inputNum1 {

		inputNum1, inputNum2 = inputNum2, inputNum1
	}

	var i int

	for i = inputNum1; i <= inputNum2; i++ {

		if i%2 == 0 {

			fmt.Printf("Even numbers : %d\n", i)
		}

		if i%2 != 0 {

			defer fmt.Printf("Odd numbers : %d\n", i)
		}

	}

}
