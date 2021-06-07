package main

import "fmt"

func searchInsert(nums []int, target int) int {
	//if target is small or equal to element in array, index will be return
	//if target is bigger than element in array index will increment by 1 and will get return
	biggerTarget := 0
	for i, v := range nums {
		if v >= target {
			return i
		}
		biggerTarget++
	}
	return biggerTarget
}

func searchInsert1(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	mid := 0
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	if nums[mid] > target {
		return mid
	}
	return mid + 1
}

func searchInsert2(nums []int, target int) int {
	l := len(nums)

	s, e := 0, l-1
	mid := (s + e) / 2

	for s <= e {
		switch true {
		case (nums[mid] == target):
			return mid
		case (target < nums[mid]): // search  lower part of slice
			e = mid - 1
			mid = (s + e) / 2
		case (target > nums[mid]): // search upper part of slice
			s = mid + 1
			mid = (s + e) / 2
		}
	}

	return s
}

func main() {

	nums := []int{1, 3, 5, 6}

	ans := searchInsert1(nums, 0)

	fmt.Println(ans)
	fmt.Println(nums)

}
