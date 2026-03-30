/*
Min Size Subarray Sum
Minimum Size Subarray Sum:

Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example:

target = 15

nums = [5, 1, 3, 5, 10, 7, 4, 9, 2, 8]

expected = 2
*/

package main

func minSubArrayLen(target int, nums []int) int {
	left := 0
	currentSum := 0
	minLen := int(^uint(0) >> 1) // max int

	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]

		for currentSum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			currentSum -= nums[left]
			left++
		}
	}

	if minLen == int(^uint(0)>>1) {
		return 0
	}

	return minLen
}
