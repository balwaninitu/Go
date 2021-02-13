package main

import "fmt"

func ptr(x int) {
	x = 5

}

func main() {

	x := 1

	ptr(x)

	fmt.Println(x)

}
