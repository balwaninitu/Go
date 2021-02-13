package main

func main() {

	var (
		harrypotter = book{title: "hungry witch", price: 24.5}
		ludo        = &game{title: "ludo", price: 20}
		tetris      = &game{title: "tetris", price: 5}
		rubik       = puzzle{title: "rubik", price: 5}
		jenga       = toy{title: "yoda", price: 150}
	)

	var store list
	store = append(store, ludo, tetris, harrypotter, rubik, &jenga)

	store.discount(.5)
	store.print()

}
