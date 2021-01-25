package main

import "fmt"

func main() {

	var (
		harrypotter = book{title: "hungry witch", price: 24.5}
		ludo        = game{title: "ludo", price: 20}
		tetris      = game{title: "tetris", price: 5}
		rubik       = puzzle{title: "rubik", price: 5}
	)

	var store list
	store = append(store, &ludo, &tetris, harrypotter, rubik)
	store.print()

	//interface values are comparable

	fmt.Println(store[0] == &ludo)
	fmt.Println(store[3] == rubik)

}
