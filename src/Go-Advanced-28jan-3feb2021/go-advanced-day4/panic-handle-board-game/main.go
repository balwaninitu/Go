package main

import (
	"fmt"
)

var n int
var games []string = []string{"Ludo", "Chess", "puzzle"}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error")
		}
	}()

	var i int
	var v string
	for i, v = range games {
		i = i + 1
		fmt.Printf("%d: %s\n", i, v)
	}

	printing()

}

func printing() {

	fmt.Println("The Board games are:")
	fmt.Println()
	fmt.Println("Whats is your favourite game")
	fmt.Scanln(&n)
	fmt.Printf("Oh I see, your fav game is: %v\n", games[n-1])

}
