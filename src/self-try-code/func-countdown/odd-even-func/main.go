package main

import "fmt"

func oddEven(n int) (int, bool) {
	if n%2 == 0 {

		return n, true
	}

	return n, false
}

func main() {

	fmt.Println(oddEven(5))

}
