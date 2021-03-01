package main

import (
	"fmt"
	"math"
)

type errorMsg struct {
	Msg  string
	Line int
	Char int
}

func (e *errorMsg) Error() string {
	return e.Msg
}

func calCircleArea(r float64) (float64, error) {

	if r < 0.0 {
		//		return 0, errors.New("Invalid value")
		return 0, &errorMsg{"Enter float value", 2, 1}
	}

	return math.Pi * r * r, nil

}
func main() {

	area, err := calCircleArea(-1)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("Area of circle is %.2f", area)

	}
}
