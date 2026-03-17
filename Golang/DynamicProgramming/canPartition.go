package main

func canPartition(nums []int) bool {
	total := 0
	for _, v := range nums {
		total += v
	}

	if total%2 != 0 {
		return false
	}

	target := total / 2
	n := len(nums)

	prev := make([]bool, target+1)
	prev[0] = true

	for i := 1; i <= n; i++ {
		curr := make([]bool, target+1)
		curr[0] = true

		for j := 1; j <= target; j++ {
			if nums[i-1] <= j {
				curr[j] = prev[j-nums[i-1]]
			}

			curr[j] = curr[j] || prev[j]
		}

		prev = curr
	}

	return prev[target]
}
