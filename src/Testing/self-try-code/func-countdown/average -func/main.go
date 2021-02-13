package main

import "fmt"

func average(slice []int) (ave float64) {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	ave = (float64(sum / len(slice)))
	return
}

func main() {

	slice1 := []int{5, 6, 2, 1, 9, 80}

	fmt.Println(average(slice1))

}
