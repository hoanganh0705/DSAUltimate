package main

func unboundedKnapsack(N int, W int, val []int, wt []int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {

			exclude := dp[i-1][j]

			include := 0
			if wt[i-1] <= j {
				include = val[i-1] + dp[i][j-wt[i-1]]
			}

			if include > exclude {
				dp[i][j] = include
			} else {
				dp[i][j] = exclude
			}
		}
	}

	return dp[N][W]
}
