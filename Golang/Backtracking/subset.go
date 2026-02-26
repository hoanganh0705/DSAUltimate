package main

func powerSet(nums []int) [][]int {
	var output [][]int

	// Declare the helper function signature first so it can call itself recursively
	var helper func(i int, subset []int)

	helper = func(i int, subset []int) {
		if i == len(nums) {
			// Equivalent to subset.copy()
			subsetCopy := make([]int, len(subset))
			copy(subsetCopy, subset)
			output = append(output, subsetCopy)
			return
		}

		// Exclude the current element
		helper(i+1, subset)

		// Include the current element
		subset = append(subset, nums[i])
		helper(i+1, subset)

		// Backtrack (Equivalent to subset.pop())
		subset = subset[:len(subset)-1]
	}

	helper(0, []int{})
	return output
}
