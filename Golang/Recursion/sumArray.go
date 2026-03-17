// Sum of array elements
//
// No loops
//
// Example: [1, 2, 3, 4] → 10

package main

func sumPtr(arr *[]int, i int) int {
	if i >= len(*arr) {
		return 0
	}
	return (*arr)[i] + sumPtr(arr, i+1)
}
