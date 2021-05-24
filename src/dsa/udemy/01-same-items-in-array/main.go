package main

import "fmt"

func unorderedEqual(first, second []string) bool {
	// if len(first) != len(second) {
	//     return false
	// }
	exists := make(map[string]bool)
	for _, value := range first {
		//all the values in the map = true
		exists[value] = true
	}
	//fmt.Println("map", exists)
	for _, value := range second {
		if !exists[value] {
			//return false
		}
	}
	return false
}

func main() {
	//array1 := []string{"a", "b", "x", "z"}

	//array2 := []string{"i", "v", "w", "a"}

	first := []string{"hi", "there", "stack", "overflow"}
	second := []string{"stack", "overflow", "hi", "there"}
	fmt.Println(unorderedEqual(first, second))

	first = []string{"hello", "there", "stack", "overflow"}
	second = []string{"s", "over", "hi", "the"}
	fmt.Println(unorderedEqual(first, second))

}
