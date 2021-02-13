package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Println("enter value 1")
	fmt.Scanln(&num1)

	fmt.Println("enter value 1")
	fmt.Scanln(&num2)

	if num1 > num2 {

		num1, num2 = num2, num1
	}

	if num1 < num2 {

		for i := num1; i <= num2; i += 2 {

			fmt.Println("count up", i)

		}
		fmt.Println()

		for i := num2; i >= num1; i -= 2 {

			fmt.Println("count up", i)

		}

	} else if num1 > num2 {

		for i := num1; i >= num2; i -= 2 {

			fmt.Println("count down", i)
		}

	}

}
