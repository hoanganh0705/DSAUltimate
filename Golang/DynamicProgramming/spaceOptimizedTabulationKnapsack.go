package main

func spaceOptimizedTabulationKnapsack(W int, wt []int, val []int, n int) int {
	prev := make([]int, W+1)
	curr := make([]int, W+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= W; j++ {

			exclude := prev[j]

			include := 0
			if wt[i-1] <= j {
				include = val[i-1] + prev[j-wt[i-1]]
			}

			if include > exclude {
				curr[j] = include
			} else {
				curr[j] = exclude
			}
		}

		// copy curr into prev
		copy(prev, curr)
	}

	return prev[W]
}
