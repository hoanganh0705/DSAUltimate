/*
Determine how to optimally fill a knapsack with a capacity of W kilograms using a list of N items, where each item is represented by a pair [profit, weight].
In the Fractional Knapsack problem, you can take fractions of items to maximize the total profit in the knapsack.(N will be greater than equal to 1 )

Example 1:

Given arr[] = [[70, 10], [90, 20], [150, 30]]

W= 25

Expected output = 145

Example 2:

Given arr[] = [[70, 10], [90, 20], [150, 30]]

W= 45

Expected output = 242.5

*/

package main

import "sort"

func fractionalKnapsack(W int, arr [][]int, n int) float64 {
	// sort by value/weight ratio (descending)
	sort.Slice(arr, func(i, j int) bool {
		return float64(arr[i][0])/float64(arr[i][1]) > float64(arr[j][0])/float64(arr[j][1])
	})

	remainingWeight := W
	value := 0.0
	n = len(arr)

	for i := 0; i < n; i++ {
		if remainingWeight == 0 {
			break
		}

		weight := min(remainingWeight, arr[i][1])
		remainingWeight -= weight
		value += (float64(arr[i][0]) / float64(arr[i][1])) * float64(weight)
	}

	return value
}
