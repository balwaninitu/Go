package main

import "fmt"

func main() {

	num1 := 17

	num2 := 24

	if num1 > num2 {

		fmt.Printf("%d is bigger than %d\nIt is bigger by %d", num1, num2, (num1 - num2))
	}

	if num2 > num1 {

		fmt.Printf("%d is bigger than %d\nIt is bigger by %d", num2, num1, (num2 - num1))

	}

}
