package main

import (
	"fmt"
	"math"
)

func calSqrt(a float64) (result float64, ok bool) {
	if a >= 0 {
		result = math.Sqrt(a)
		ok = true
	}
	return result, ok
}

//alternate type of same func

func altCalSqrt(a float64) (result float64, ok bool) {
	if a < 0 {
		return
	}
	return math.Sqrt(result), true
}

func main() {
	//add blan identifier if seconf=d value dont need
	var ans, _ = calSqrt(64)

	fmt.Printf("Result is %.2f \n", ans)

}
