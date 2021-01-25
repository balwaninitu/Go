package main

import "fmt"

func main() {

	store := list{
		&book{product{"hungry witch", 24.5}, 118281600},
		&book{product{"evil witch", 24.5}, "733622400"},
		&book{product{"petit witch", 24.5}, nil},
		&game{product{"ludo", 20}},
		&game{product{"tetris", 5}},
		&puzzle{product{"rubik", 5}},
		&toy{product{"yoda", 150}},
	}

	store.discount(.5)
	store.print()

	t := &toy{product{"yoda", 150}}
	fmt.Printf("%#v\n", t)

	b := &book{product{"hungry witch", 24.5}, 118281600}
	fmt.Printf("%#v\n", b)
}
