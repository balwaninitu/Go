/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Example 2:

Input: nums = [1]
Output: 1
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23

*/

package main

import "fmt"

func maxSubArray1(nums []int) int {
	return divide(nums, 0, len(nums)-1)
}

func divide(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}

	mid := (start + end) / 2
	left := divide(nums, start, mid)
	right := divide(nums, mid+1, end)
	both := conquer(nums, start, mid, end)
	return max1(left, max1(right, both))
}

func conquer(nums []int, start, mid, end int) int {
	leftSum := nums[mid]
	rightSum := nums[mid+1]

	sum := 0
	for i := mid; i >= start; i-- {
		sum += nums[i]
		leftSum = max1(leftSum, sum)
	}

	sum = 0
	for i := mid + 1; i <= end; i++ {
		sum += nums[i]
		rightSum = max1(rightSum, sum)
	}

	return leftSum + rightSum
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// To solve this problem, we rely on updating the current value only if its value would increase by adding its previous
// one. On every iteration, we update the maximum subarray sum by checking if the current one is bigger than the
// temporary maximum.
//
// T: O(n)
// S: O(1)

//intially maxsum will be 5, first element of array
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//1st iteration
	//maxsum --> 5(n2)
	//i = 1 --> 4
	//loop starts with element 4 (index 1) where as maxsum is element 5 of array
	//condition 5 > 4 i.e. 5+4 = 9(n1)
	//max func after first iteration --> n1 = 9, n2=5
	//when n1> n2, n1 get return so maxsum will be 9

	//2nd iteration
	//n2(maxsum) = 9, i = 2(-1), n1 = 8(9-1)
	//n2 > n1 return n2, maxSum: 9

	//3rd iteration
	//maxSum: 9, i = 3(7), n1 = 15

	//4th iteration
	//maxsum: 15, i: 4(8), n1 = 23,

	//5th iteration
	//maxsum = 23, return n2
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		//int(bigger integer) = (n1, n2)
		maxSum = max(nums[i], maxSum)
	}
	return maxSum
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {

	nums := []int{5, 4, -1, 7, 8}

	ans := maxSubArray(nums)

	fmt.Println(ans)
}

// dynamic-programming solution
func maxSubArray3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// holy shit - this solution is fantastically
	// simple

	// the intuition comes from us,
	// rephrasing the problem as:

	// If f(i) the max sum with any of the
	// of the subarray ending with index i

	// Then -

	// f(i) = max(
	//	A[i] + f(i - 1),
	//	A[i]
	// )

	solutions := make([]int, len(nums))
	solutions[0] = nums[0]

	finalSolution := solutions[0]

	for i := 1; i < len(nums); i++ {
		solutions[i] = max3(
			nums[i],
			nums[i]+solutions[i-1],
		)

		if solutions[i] > finalSolution {
			finalSolution = solutions[i]
		}
	}

	return finalSolution
}

func max3(values ...int) int {
	// assume at least one input
	maxVal := values[0]

	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}
