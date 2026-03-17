package main

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)

	summation := 0
	for _, v := range nums {
		summation += v
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2*summation+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var helper func(int, int) int
	helper = func(index int, sumNums int) int {
		if index < 0 {
			if sumNums == target {
				return 1
			}
			return 0
		}

		if dp[index][sumNums+summation] != -1 {
			return dp[index][sumNums+summation]
		}

		negative := helper(index-1, sumNums-nums[index])
		positive := helper(index-1, sumNums+nums[index])

		dp[index][sumNums+summation] = negative + positive
		return dp[index][sumNums+summation]
	}

	return helper(n-1, 0)
}
