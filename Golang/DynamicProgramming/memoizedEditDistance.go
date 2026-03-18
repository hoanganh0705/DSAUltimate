package main

func memoizedMinDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var helper func(int, int) int
	helper = func(i int, j int) int {
		// base cases
		if i >= n && j >= m {
			return 0
		}
		if i >= n {
			return m - j
		}
		if j >= m {
			return n - i
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		if word1[i] == word2[j] {
			dp[i][j] = helper(i+1, j+1)
		} else {
			replace := 1 + helper(i+1, j+1)
			delete := 1 + helper(i+1, j)
			insert := 1 + helper(i, j+1)

			dp[i][j] = min(replace, delete, insert)
		}

		return dp[i][j]
	}

	return helper(0, 0)
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}
