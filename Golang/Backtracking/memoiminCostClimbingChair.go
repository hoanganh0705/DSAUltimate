package main

import "math"

func memoMinCostClimbingStairs(cost []int) int {

	n := len(cost)
	minCost := make([]int, n)

	for i := range minCost {
		minCost[i] = -1
	}

	var helper func(index int) int

	helper = func(index int) int {

		// base case
		if index > n-1 {
			return 0
		}

		if minCost[index] != -1 {
			return minCost[index]
		}

		// one step
		onestep := cost[index] + helper(index+1)

		// two steps
		twostep := cost[index] + helper(index+2)

		minCost[index] = int(math.Min(float64(onestep), float64(twostep)))

		return minCost[index]
	}

	return int(math.Min(float64(helper(0)), float64(helper(1))))

}
