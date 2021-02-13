package main

import "fmt"

func init() {
	fmt.Println("init in second.go")
}

func init() {
	fmt.Println("init again in second.go")
}

func init() {
	fmt.Println("init again & again  in second.go")
}
