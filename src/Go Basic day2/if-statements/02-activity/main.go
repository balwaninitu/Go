package main

import "fmt"

func main() {

	num := 27

	if num%2 == 0 {

		fmt.Printf("%d is even\n", num)

	} else {

		fmt.Printf("%d is odd\n", num)
	}

	if num > 10 {

		fmt.Printf("%d  have more than 1 digit\n", num)

	}

}
