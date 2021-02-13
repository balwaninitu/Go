package main

// func selectionSort(arr []int, n int) {

// }

func largestNum(array [5]int, n int) int {

	largest := 0
	for i := 1; i < len(array); i++ {
		if array[i] > array[largest] {
			largest = i

		}
	}
	return largest
}

func main() {

	// array := [5]int{3, 4, 5, 1, 2}

	// n := largestNum(array)
	// fmt.Print(n)

}
