// Find max in array
//
// Recursively compare elements
//
// No sorting, no loops

package main

func findMax(arr []int, i int) int {
	if len(arr) == 0 {
		panic("findMax: empty slice")
	}
	// base case: last element
	if i == len(arr)-1 {
		return arr[i]
	}
	// max of the rest
	m := findMax(arr, i+1)
	if arr[i] > m {
		return arr[i]
	}
	return m
}

func Max(arr []int) int {
	return findMax(arr, 0)
}
