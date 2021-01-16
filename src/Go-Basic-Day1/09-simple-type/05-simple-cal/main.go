package main

import "fmt"

func main() {

	fun1 := "+"
	fun2 := "-"
	fun3 := "*"
	fun4 := "/"

	var val1, val2 int

	var fun string

	fmt.Println("Enter value1 from 0 to 9")

	fmt.Scanln(&val1)

	fmt.Println("Enter arithmetic functions + - * /")

	fmt.Scanln(&fun)

	fmt.Println("Enter value2 from 0 to 9")

	fmt.Scanln(&val2)

	if fun == fun1 {

		fmt.Println("Add:", val1+val2)

	} else if fun == fun2 {

		fmt.Println("Sub:", val1-val2)

	} else if fun == fun3 {

		fmt.Println("Multiply:", val1*val2)

	} else if fun == fun4 {

		fmt.Println("Divide:", val1/val2)
	}

}
