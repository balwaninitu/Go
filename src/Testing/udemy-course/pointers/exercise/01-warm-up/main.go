package main

import "fmt"

type computer struct {
	brand string
}

func main() {

	ptr := &computer{}

	if ptr == nil {

		fmt.Println("ptr is nil")
	} else {

		fmt.Println("ptr is not nil", ptr)
	}

	var null *computer

	if null == nil {

		fmt.Println("null is nil")
	}

	//apple := &computer{brand : "apple"}

}
