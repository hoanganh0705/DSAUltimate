package main

func memoizedLCS(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var helper func(i int, j int) int
	helper = func(i int, j int) int {
		// base case
		if i >= n || j >= m {
			return 0
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		// match
		if text1[i] == text2[j] {
			dp[i][j] = 1 + helper(i+1, j+1)
		} else {
			// no match
			dp[i][j] = max(
				helper(i, j+1),
				helper(i+1, j),
			)
		}

		return dp[i][j]
	}

	return helper(0, 0)
}
