package main

import "fmt"

func main() {

	temperature := [24]float64{20.1, 24, 27.3, 30.1, 26.4, 22.2,
		20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3,
		30.1, 26.4, 20.1, 24, 27.3}
	var sum float64
	for _, v := range temperature {
		sum += v
	}

	fmt.Printf("avaerage temp of room %.2f\n", sum/float64(len(temperature)))

}
