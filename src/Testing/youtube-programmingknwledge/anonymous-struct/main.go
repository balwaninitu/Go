package main

import "fmt"

func main() {

	company := struct {
		name string
		age  int
	}{name: "a", age: 40}
	fmt.Println(company.name)
	fmt.Println(company.age)

}
