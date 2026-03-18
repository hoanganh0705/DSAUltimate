package main

func tabulationMinDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// base cases
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}

	// fill dp
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {

			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				replace := 1 + dp[i-1][j-1]
				delete := 1 + dp[i-1][j]
				insert := 1 + dp[i][j-1]

				dp[i][j] = min2(replace, delete, insert)
			}
		}
	}

	return dp[n][m]
}

func min2(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}
