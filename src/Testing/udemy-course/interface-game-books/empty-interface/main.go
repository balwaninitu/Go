package main

import "fmt"

func main() {
	var any interface{}

	any = []int{1, 2, 3}
	any = map[int]bool{1: true, 2: false}
	any = "hello"
	any = 3
	//cant do below
	//any = any*3

	//istead use assertion method any.(type)*value
	any = any.(int) * 2
	fmt.Println(any)

}
