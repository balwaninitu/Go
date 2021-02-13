package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var number1 string
		fmt.Println("Enter the first number: ")
		fmt.Scanln(&number1)
		var number2 string
		fmt.Println("Enter the second number: ")
		fmt.Scanln(&number2)
		number1Value, _ := strconv.ParseInt(number1, 10, 0)
		number2Value, _ := strconv.ParseInt(number2, 10, 0)
		if number1Value < number2Value {
			for i := number1Value; i <= number2Value; i++ {
				if i%2 != 0 {
					fmt.Println("The number is odd!", i)
				}
			}
			for i := number1Value; i <= number2Value; i++ {
				if i%2 == 0 {
					fmt.Println("The number is even!", i)
				}
			}
		}
		if number1Value > number2Value {
			for i := number2Value; i <= number1Value; i++ {
				if i%2 != 0 {
					fmt.Println("The number is odd!", i)
				}
			}
			for i := number1Value; i <= number2Value; i++ {
				if i%2 == 0 {
					fmt.Println("The number is even!", i)
				}
			}
		}
	}
}
