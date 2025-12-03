/*

Problem:
Given an array of positive integers nums and a target T,
return the number of ways to pick elements (each element can be used multiple times) so they sum to exactly T.

Example:
nums = [1, 3, 4], T = 5
Possible ways:

1+1+1+1+1

1+1+3

1+4

3+1+1
(etc… order matters or doesn’t matter—you pick, but be consistent)

*/

package main

import "fmt"

func countWays(nums []int, target int, i int) int {
	// Base case: exact match → 1 valid way
	if target == 0 {
		return 1
	}

	// Base case: out of bounds or overshoot → no valid way
	if target < 0 || i == len(nums) {
		return 0
	}

	// Option 1: take nums[i] (stay at same index because reuse allowed)
	take := countWays(nums, target-nums[i], i)

	// Option 2: skip nums[i] (move to next index)
	skip := countWays(nums, target, i+1)

	return take + skip
}

func main() {
	nums := []int{1, 3, 4}
	target := 5
	fmt.Println(countWays(nums, target, 0))
}
