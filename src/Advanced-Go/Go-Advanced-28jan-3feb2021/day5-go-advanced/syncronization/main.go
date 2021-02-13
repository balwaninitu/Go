package main

import "fmt"

func main() {

}

func addGame(name string, price int) {
	var gameName string
	var priceGame int

	fmt.Println("name of game to add")
	fmt.Scanln(&gameName)
	fmt.Println("price of game to add")
	fmt.Scanln(&priceGame)
	m := make(map[string]int)
	m[gameName] = priceGame

}
