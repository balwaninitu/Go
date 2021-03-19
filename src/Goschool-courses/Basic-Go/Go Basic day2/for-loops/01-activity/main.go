package main

import "fmt"

func main() {

	var i int

	for i = 0; i < 1000; i++ {

		fmt.Println("Numbers", i)
	}

	for i = 0; i <= 1000; i++ {

		if i%2 == 0 {

			fmt.Println("Even numbers", i)

		}

	}

	for i = 0; i <= 1000; i++ {

		if i%2 != 0 {

			fmt.Println("Odd numbers", i)

		}

	}
}
