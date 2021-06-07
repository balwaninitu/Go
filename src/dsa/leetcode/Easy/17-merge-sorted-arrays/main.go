package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	nums1 = []int{1, 2, 3, 0, 0, 0}

	m = len(nums1)

	nums2 = []int{2, 5, 6}

	n = len(nums2)

}

func main() {

	nums1 := []int{1, 2, 3}

	m := len(nums1)

	fmt.Println(m)

	nums2 := []int{2, 5, 6}

	n := len(nums2)
	fmt.Println(n)

	for i, v1 := range nums1 {
		for _, v2 := range nums2 {
			if v1 < v2 {
				v2 = nums1[i+1]
			} else if v1 > v2 {
				v2 = nums1[i-1]

			}
		}
	}
	fmt.Println("nums1", nums1)

}
