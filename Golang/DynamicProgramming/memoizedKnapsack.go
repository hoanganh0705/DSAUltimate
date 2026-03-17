package main

func memoizedKnapsack(W int, wt []int, val []int, n int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, W+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var helper func(index, remWeight int) int
	helper = func(index, remWeight int) int {
		if index >= n || remWeight == 0 {
			return 0
		}

		if dp[index][remWeight] != -1 {
			return dp[index][remWeight]
		}

		exclude := helper(index+1, remWeight)
		include := 0

		if wt[index] <= remWeight {
			include = val[index] + helper(index+1, remWeight-wt[index])
		}

		dp[index][remWeight] = max(exclude, include)
		return dp[index][remWeight]
	}

	return helper(0, W)
}
