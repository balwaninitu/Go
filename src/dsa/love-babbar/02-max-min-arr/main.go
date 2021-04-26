package main

import "fmt"

func main() {

	arr := []int{9, 6, 7, 56, 23, 78}

	//min := arr[0]
	//max := arr[1]
	// fmt.Println(min)
	// fmt.Println(arr[1])
	// fmt.Println(arr[2])
	// if min > arr[1] {
	// 	min, arr[1] = arr[1], min
	// }
	// if min > arr[2] {
	// 	min, arr[2] = arr[1], min
	// }

	// fmt.Println("=======After======")
	// fmt.Println(min)
	// fmt.Println(arr[1])
	// fmt.Println(arr[2])

	//instead of hardcore as above add variable

	var n, min, max int

	for _, v := range arr {
		if v > n {
			//fmt.Println(v, ">", n)
			n = v
			max = n
		} else {
			//fmt.Println(v, "<", n)
		}
	}
	fmt.Println("max num", max)

	for _, v := range arr {
		if v > n {
			//fmt.Println(v, ">", n)
		} else {
			//fmt.Println(v, "<", n)
			n = v
			min = v
		}
	}
	fmt.Println("min num", min)
}
