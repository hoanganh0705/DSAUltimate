/*
Coding Exercise ( Permutations)
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:

nums = [1,4]

Output :[[1,4],[4,1]]

Example 2:

nums = [1,4,5]

Output: [[1,4,5],[1,5,4],[4,1,5],[4,5,1],[5,1,4],[5,4,1]]
*/
package main

func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums)-1 {
			// Make a copy of nums and append to result
			perm := make([]int, len(nums))
			copy(perm, nums)
			result = append(result, perm)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start] // Swap
			backtrack(start + 1)                        // Recurse
			nums[start], nums[i] = nums[i], nums[start] // Backtrack (swap back)
		}
	}
	backtrack(0)
	return result
}
