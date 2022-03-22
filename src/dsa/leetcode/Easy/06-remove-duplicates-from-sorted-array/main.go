/*
Given a sorted array nums, remove the duplicates in-place such that 
each element appears only once and returns the new length.

Do not allocate extra space for another array, you must do this by modifying the 
input array in-place with O(1) extra memory.

Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, 
which means a modification to the input array will be known to the caller as well.

Internally you can think of this:

// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}


Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2]
Explanation: Your function should return length = 2, with the first 
two elements of nums being 1 and 2 respectively. 
It doesn't matter what you leave beyond the returned length.

Example 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4]
Explanation: Your function should return length = 5, 
with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively. 
It doesn't matter what values are set beyond the returned length.
*/

package main

import "fmt"

//Given a sorted array nums, remove the duplicates in-place such that
//each element appears only once and returns the new length.

//return slice instead of length
func removeDuplicates(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums[0:n]
	}
	j := 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return nums[0:j]
}

/*Since the array is already sorted, we can keep two pointers i and j,
where i is the slow-runner while j is the fast-runner.
As long as nums[i] = nums[j], we increment j to skip the duplicate.
When we encounter nums[j] not equal to nums[i]
the duplicate run has ended so we must copy its value to nums[i + 1].
i is then incremented and we repeat the same process again until j reaches the end of array.
*/
func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	i := 0
	for j := 1; j < n; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	//here i +1 because i starts from 0 whereas j is 1
	return i + 1
}

/*Solution in mind
Make use of 2 pointers, a fast and slow moving pointer.
Initialise fast to 1 and slow to 0.While fast has not reached the end of array,
check if the elements at index fast and slow are same, if yes, move fast forward by 1 place.
If not, copy the non duplicate element into the position after slow.
Once fast reaches the end of array,
all unique elements will be stored inside the first "slow" number of elements in the array.
*/
func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}
	slow := 0
	fast := 1
	n := len(nums)

	for fast < n {
		if nums[fast] == nums[slow] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
func main() {
	//fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates1([]int{1, 2}))
	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 4, 5, 5, 6, 7, 8, 8, 8, 9}))
}
