package main

import "fmt"

func main() {

	var year int

	fmt.Println("Please enter year")

	fmt.Scanln(&year)

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {

		fmt.Println("Given year is leap year")
	} else {

		fmt.Println("Given year is not leap year")
	}

}
