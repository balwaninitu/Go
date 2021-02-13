package main

import "fmt"

func main() {

	var recurse func()

	recurse = func() {
		fmt.Println("welcome")
		recurse()

	}
	recurse()
}
