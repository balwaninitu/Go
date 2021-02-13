package main

func main() {

	l := list{
		{title: "hungry witch", price: 24.5, released: toTimestamp(118281600)},
		{title: "evil witch", price: 24.5, released: toTimestamp("733622400")},
		{title: "petit witch", price: 24.5},
		{title: "ludo", price: 20},
		{title: "tetris", price: 5},
		{title: "rubik", price: 5},
		{title: "yoda", price: 150},
	}

	l.discount(.5)
	l.print()

}
