package main

import (
	"errors"
	"fmt"
)

type invalidTriangleError struct {
	a, b, c int
	msg     string
}

// func (i *invalidTriangleError) error() string {
// 	return "side is less than sum of other two sides"

// }

func createTriangle(a, b, c int) (int, error) {
	if a < 0 || b < 0 || c < 0 {
		return 0, errors.New("side is less than sum of other two sides")
	}

	return (a + b + c/2), nil
}

func main() {

	area, err := createTriangle(12, 5, 20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("arear of tringle %d\n", area)

	}

}
