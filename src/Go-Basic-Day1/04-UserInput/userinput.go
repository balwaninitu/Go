package main

import "fmt"

func main() {

	num := 23 //predetermined number

	var guess int // number input by user

	fmt.Println("Enter an integer value:  ")

	fmt.Scanln(&guess)

	if num == guess {

		fmt.Println("Well Done! Your guess is correct")
	} else if guess < num {

		fmt.Println("Too low, try again next time!")

	} else if guess > num {

		fmt.Println("Too high, try again next time!”")
	}

}
