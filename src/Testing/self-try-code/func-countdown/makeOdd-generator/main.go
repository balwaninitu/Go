package main

import "fmt"

func makeOddgenerator() func() uint {
	i := uint(1)
	return func() (num uint) {
		num = i
		i += 2
		return
	}
}
func main() {

	nextOdd := makeOddgenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
