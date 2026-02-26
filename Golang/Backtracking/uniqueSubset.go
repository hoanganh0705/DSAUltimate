package main

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	// The array must be sorted for the duplicate-skipping logic to work
	sort.Ints(nums)
	var res [][]int

	var helper func(i int, curr []int)
	helper = func(i int, curr []int) {
		// base
		if i == len(nums) {
			currCopy := make([]int, len(curr))
			copy(currCopy, curr)
			res = append(res, currCopy)
			return
		}

		// include
		curr = append(curr, nums[i])
		helper(i+1, curr)

		// backtracking (equivalent to curr.pop())
		curr = curr[:len(curr)-1]

		// exclude
		// Skip duplicate elements to avoid duplicate subsets
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}

		// Included the missing 'curr' argument here!
		helper(i+1, curr)
	}

	helper(0, []int{})
	return res
}
