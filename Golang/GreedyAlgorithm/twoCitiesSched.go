/*
A company is planning to interview 2n people. Given the array costs where costs[i] = [aCosti, bCosti], the cost of flying the ith person to city a is aCosti, and the cost of flying the ith person to city b is bCosti.

Return the minimum cost to fly every person to a city such that exactly n people arrive in each city.
*/

package main

import (
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})

	n := len(costs)
	cost := 0

	for i := 0; i < n/2; i++ {
		cost += costs[i][0]
	}

	for i := n / 2; i < n; i++ {
		cost += costs[i][1]
	}

	return cost
}
