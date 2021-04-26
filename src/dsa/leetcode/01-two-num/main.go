package main

import "fmt"

func twonum(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{15, 8, 7, 2}

	target := 9

	ans := twonum(nums, target)

	fmt.Println(ans)
}
