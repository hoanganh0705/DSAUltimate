package main

func knapSack(W int, wt []int, val []int, n int) int {
	// dp[i][j] = max value using first i items with capacity j
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= W; j++ {

			exclude := dp[i-1][j]

			include := 0
			if wt[i-1] <= j {
				include = val[i-1] + dp[i-1][j-wt[i-1]]
			}

			if include > exclude {
				dp[i][j] = include
			} else {
				dp[i][j] = exclude
			}
		}
	}

	return dp[n][W]
}
