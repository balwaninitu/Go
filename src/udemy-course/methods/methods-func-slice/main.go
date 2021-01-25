package main

func main() {

	var (
		//harrypotter := book{title: "hungry witch", price: 24.5}
		ludo   = game{title: "ludo", price: 20}
		tetris = game{title: "tetris", price: 5}
	)

	var items []*game
	items = append(items, &ludo, &tetris)

	my := list(items)
	//my = nil
	my.print()
}
