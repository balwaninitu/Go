package main

import (
	"fmt"
	"time"
)

func init() {
	time.Sleep(100 * time.Millisecond)

	fmt.Println("init func")
}

func init() {
	fmt.Println("this is second init")
}

func init() {
	fmt.Println("init third time in main.go")
}

func main() {

	fmt.Println("main func")

}
