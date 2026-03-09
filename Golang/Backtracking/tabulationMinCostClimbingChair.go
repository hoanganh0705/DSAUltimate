package main

import "math"

func tabulationMinCostClimbingStairs(cost []int) int {

	n := len(cost)

	mincost := make([]int, n+1)

	mincost[0] = 0
	mincost[1] = 0

	for i := 2; i <= n; i++ {

		onestep := cost[i-1] + mincost[i-1]
		twostep := cost[i-2] + mincost[i-2]

		mincost[i] = int(math.Min(float64(onestep), float64(twostep)))
	}

	return mincost[n]
}
