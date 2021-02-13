package main

import "fmt"

func main() {

	num := 35

	p := &num

	v := *p

	fmt.Println(v)

	fmt.Println(*p)

}
