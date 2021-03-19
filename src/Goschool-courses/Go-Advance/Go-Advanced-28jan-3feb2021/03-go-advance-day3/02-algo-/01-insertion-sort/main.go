package main

import "fmt"

func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {

		data := arr[i]

		last := i
		for (last > 0) && (arr[last-1] > data) {
			arr[last] = arr[last-1]
			last--
		}
		arr[last] = data
	}
}

func main() {

	array := []int{20, 80, 40, 25, 60, 40}

	insertionSort(array, len(array))
	fmt.Println(array)

}
