/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?



Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Approach 1: Brute Force
Algorithm

In this brute force approach we take all possible step combinations i.e. 1 and 2, at every step. At every step we are calling the function climbStairsclimbStairs for step 11 and 22, and return the sum of returned values of both functions.

climbStairs(i,n)=(i + 1, n) + climbStairs(i + 2, n)climbStairs(i,n)=(i+1,n)+climbStairs(i+2,n)

where ii defines the current step and nn defines the destination step.

Complexity Analysis

Time complexity : O(2^n)O(2
n
 ). Size of recursion tree will be 2^n2
n
 .

Recursion tree for n=5 would be like this:

Climbing_Stairs

Space complexity : O(n)O(n). The depth of the recursion tree can go upto nn.


Approach 2: Recursion with Memoization
Algorithm

In the previous approach we are redundantly calculating the result for every step. Instead, we can store the result at each step in memomemo array and directly returning the result from the memo array whenever that function is called again.

In this way we are pruning recursion tree with the help of memomemo array and reducing the size of recursion tree upto nn.

*/

package main

import "fmt"

func climbStairs(n int) int {

	cache := map[int]int{}

	if n == 0 {
		return 1
	} else if n < 3 {
		return n
	}

	if output, ok := cache[n]; ok {
		return output

	}
	cache[n] = climbStairs(n-2) + climbStairs(n-1)
	return cache[n]
}

func main() {

	n := 5

	ans := climbStairs(n)

	fmt.Println(ans)
}
