package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("main fuction started")

	go func() {
		fmt.Println("Hello world")
	}()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main function ended")

}
