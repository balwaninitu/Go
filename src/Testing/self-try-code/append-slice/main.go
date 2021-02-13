package main

import "fmt"

func main() {

	slice := []int{0, 4, 3}

	s := append(slice, 9)

	fmt.Println(slice)

	fmt.Println(s)

	slice1 := []int{}

	c := copy(slice1, slice)

	fmt.Println(c)

}
