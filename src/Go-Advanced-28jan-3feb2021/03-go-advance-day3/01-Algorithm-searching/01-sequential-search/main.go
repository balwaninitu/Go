package main

import "fmt"

func makeSlice(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func sequentialSearch(slice []int, n int, target int) (int, count int) {
	count = 0
	for i := range slice {
		count++
		if slice[i] == target {
			return i, count
		}
	}

	return -1, count
}

func binarySearch(slice []int, n int, target int) (int, count int) {
	first := 0
	last := n
	count = 0
	for first <= last {
		mid := (first + last) / 2
		if slice[mid] == target {
			return mid, count
		} else {
			if target < slice[mid] {
				last = mid - 1
			} else {
				first = mid + 1
			}
		}
	}

	return -1, count
}

func main() {

	slice := makeSlice(1, 1000)
	//fmt.Println(slice)

	IndexFound, count := sequentialSearch(slice, len(slice), 50)
	fmt.Printf("Sequential comparison %d\n", count)
	fmt.Printf("Target found at index %d\n", IndexFound)
	IndexFound, count = binarySearch(slice, len(slice), 50)
	fmt.Printf("Binary system comparison %d\n", count)
	fmt.Printf("Target found at index %d\n", IndexFound)

}
