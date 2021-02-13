package main

func main() {

	store := list{
		book{title: "hungry witch", price: 24.5, published: 118281600},
		book{title: "evil witch", price: 24.5, published: "733622400"},
		book{title: "petit witch", price: 24.5},
		&game{title: "ludo", price: 20},
		&game{title: "tetris", price: 5},
		puzzle{title: "rubik", price: 5},
		&toy{title: "yoda", price: 150},
	}

	store.discount(.5)
	store.print()

}
