package main

func main() {

	harrypotter := book{title: "hungry witch", price: 24.5}

	harrypotter.print()

	ludos := game{title: "ludo", price: 20}

	ludos.print()
	ludos.discount(0.5)

}
