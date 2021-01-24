package main

import "fmt"

func main() {

	fmt.Println("Maps")

	maps()

}

func maps() {

	confused := map[string]int{"one": 2, "two": 1}

	fix(confused)
	fmt.Println(confused)

}

func fix(m map[string]int) {

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}
