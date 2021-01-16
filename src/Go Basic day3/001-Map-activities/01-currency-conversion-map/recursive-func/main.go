package main

import (
	"fmt"
)

func factorial(factor int) int {

	if factor == 0 {

		return 1

	}

	return factor * factorial(factor-1)

}

func main() {

	var factorVal int

	fmt.Println("Enter Factorial")

	fmt.Scanln(&factorVal)

	//factorValue, _ := strconv.ParseInt(factorVal, 10, 0)

	fmt.Println(factorial(factorVal))

}
