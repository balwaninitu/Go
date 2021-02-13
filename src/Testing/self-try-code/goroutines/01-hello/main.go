package main

import (
	"fmt"
	"time"
)

func helloworld() {
	time.Sleep(9 * time.Millisecond)
	fmt.Println("Hello world")
}

func main() {

	fmt.Println("Main function started")

	go helloworld()

	time.Sleep(10 * time.Millisecond)
	fmt.Println("Main function ended")

}
