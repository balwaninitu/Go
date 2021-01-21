package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	hiddeNumber := rand.Intn(100)

	i := 1
	for i <= 5 {

		var number int

		fmt.Print("Enter Number: ")
		fmt.Scanln(&number)

		if hiddeNumber == number {

			fmt.Println("You won!")
			break
		} else if hiddeNumber == 101 {

			fmt.Println("You give up the hidden number is", hiddeNumber)
			break
		} else if number > hiddeNumber {
			fmt.Println("Too high")
		} else if number < hiddeNumber {
			fmt.Println("too low")

		}
		if i == 5 {
			fmt.Println("Game over, the correct answer is", hiddeNumber)
			break
		}

		i++

	}

}
