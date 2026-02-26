/*
	Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.
*/

package main

func permuteUnique(nums []int) [][]int {
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
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			nums[start], nums[i] = nums[i], nums[start] // Swap
			backtrack(start + 1)                        // Recurse, this will include a branch with 2 more call
			nums[start], nums[i] = nums[i], nums[start] // Backtrack (swap back)
		}
	}
	backtrack(0)
	return result
}
