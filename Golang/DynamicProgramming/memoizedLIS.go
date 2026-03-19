package main

func memoizedLIS(nums []int) int {
	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var helper func(int, int) int
	helper = func(curr int, prev int) int {
		if curr >= n {
			return 0
		}

		if dp[curr][prev+1] != -1 {
			return dp[curr][prev+1]
		}

		exclude := helper(curr+1, prev)

		include := 0
		if prev == -1 || nums[curr] > nums[prev] {
			include = 1 + helper(curr+1, curr)
		}

		if include > exclude {
			dp[curr][prev+1] = include
		} else {
			dp[curr][prev+1] = exclude
		}

		return dp[curr][prev+1]
	}

	return helper(0, -1)
}
