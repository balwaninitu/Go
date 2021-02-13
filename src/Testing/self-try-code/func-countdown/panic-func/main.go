package main

import "fmt"

func main() {

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")

	// array := [3]int{1, 2, 3}

	// s := array[4]

	// fmt.Println(s)

	// 	m := map[string]int{}

	// 	fmt.Println(m["string"])

	// }
}
