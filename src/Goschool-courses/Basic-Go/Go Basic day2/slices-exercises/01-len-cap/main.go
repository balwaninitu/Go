package main

import "fmt"

func main() {

	monthSpending := []float64{9.50, 8.00, 10.20, 7.40}

	fmt.Printf("len: %d cap: %d\n", len(monthSpending), cap(monthSpending))

	fmt.Printf("spending for week 3 is %.2f\n", monthSpending[2])

	monthSpending = append(monthSpending, 8.40, 9.40, 7.20)
	fmt.Printf("len: %d cap: %d\n", len(monthSpending), cap(monthSpending))

	subSlice := monthSpending[3:]

	fmt.Printf("len: %d cap: %d\n", len(subSlice), cap(subSlice))

}
