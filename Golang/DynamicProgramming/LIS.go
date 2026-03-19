package main

func lengthOfLIS(nums []int) int {
	n := len(nums)

	var helper func(int, int) int
	helper = func(currIndex int, prevIndex int) int {
		// base case
		if currIndex >= n {
			return 0
		}

		// exclude
		exclude := helper(currIndex+1, prevIndex)

		// include
		include := 0
		if prevIndex == -1 || nums[currIndex] > nums[prevIndex] {
			include = 1 + helper(currIndex+1, currIndex)
		}

		if include > exclude {
			return include
		}
		return exclude
	}

	return helper(0, -1)
}
