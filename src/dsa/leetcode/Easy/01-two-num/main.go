/*Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]*/

package main

import "fmt"

//bubble sort method
//brut force
//time complexity:0(n^2) --> Nested loop
func twonum(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				//fmt.Println("nums[j]", nums[j])
				return []int{i, j}
			}
		}
	}
	return nil
}

//one loop
//time complexity:0(n) --> One loop
//It turns out we can do it in one-pass.
//While we iterate and inserting elements into the table,
//we also look back to check if current element's complement already exists in the table.
//If it exists, we have found a solution and return immediately.
func twonum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		//checking if value in array is equal to target minus other value in array
		//if target = value1 + value2 then return index of those values
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		//value in array map with the index during looping
		m[v] = i
	}
	return nil
}

func main() {
	nums := []int{15, 8, 7, 2}

	target := 17

	ans := twonum(nums, target)
	ans1 := twonum1(nums, target)

	fmt.Println("nested loop", ans)
	fmt.Println("one loop", ans1)
}
