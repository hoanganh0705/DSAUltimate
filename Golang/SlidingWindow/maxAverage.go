package main

func findMaxAverage(nums []int, k int) float64 {
	currSum := 0

	for i := 0; i < k; i++ {
		currSum += nums[i]
	}

	maxSum := currSum

	for i := k; i < len(nums); i++ {
		currSum = currSum + nums[i] - nums[i-k]
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return float64(maxSum) / float64(k)
}
