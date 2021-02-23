package main

import "fmt"

func calculate(height, base float64) (area, perimeter float64) {

	area = 0.5 * base * height
	perimeter = 3 * base
	return
}

func main() {

	base, height := 10.0, 5.4

	area, perimeter := calculate(base, height)

	fmt.Printf("area: %.2f perimeter: %.2f\n", area, perimeter)

}
