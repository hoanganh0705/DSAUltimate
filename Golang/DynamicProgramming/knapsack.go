/*
You are provided with a set of N items, each with a specified weight and value.
Your objective is to pack these items into a backpack with a weight limit of W, maximizing the total value of items in the backpack. Specifically, you have two arrays: val[0..N-1], representing the values of the items, and wt[0..N-1], indicating their weights. Additionally, you have a weight limit W for the backpack.
The challenge is to determine the most valuable combination of items where the total weight does not exceed W. Note that each item is unique and indivisible, meaning it must be either taken as a whole or left entirely.



Input:
N = 3
W = 8
values[] = [2,3,9]
weight[] = [8,2,5]
Output: 12
Explanation: Choose the last 2 items that weighs 2 and 5 units respectively and hold values 3 and 9 that add up to 12.

*/

package main

func knapsack(W int, wt []int, val []int, n int) int {
	return knapsackHelper(0, W, wt, val, n)
}

func knapsackHelper(index int, remWeight int, wt []int, val []int, n int) int {
	// base case
	if index >= n || remWeight == 0 {
		return 0
	}

	// exclude current item
	exclude := knapsackHelper(index+1, remWeight, wt, val, n)

	// include current item (if possible)
	include := 0
	if wt[index] <= remWeight {
		include = val[index] + knapsackHelper(index+1, remWeight-wt[index], wt, val, n)
	}

	// choose best
	if exclude > include {
		return exclude
	}
	return include
}
