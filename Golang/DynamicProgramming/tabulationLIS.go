package main

func tabulationLIS(nums []int) int {
	n := len(nums)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {

			exclude := dp[i+1][j]

			include := 0
			if j == 0 || nums[i] > nums[j-1] {
				include = 1 + dp[i+1][i+1]
			}

			if include > exclude {
				dp[i][j] = include
			} else {
				dp[i][j] = exclude
			}
		}
	}

	return dp[0][0]
}
