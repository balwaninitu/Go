package main

import (
	"fmt"
	"modules/mypackages/numbers"
)

func main() {

	var x uint = 4
	fmt.Printf("%d is even: %t\n", x, numbers.Even(x))

}
