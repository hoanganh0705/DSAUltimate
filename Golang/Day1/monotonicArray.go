/*

Coding Exercise: Monotonic Array
Question

An array is monotonic if it is either monotone increasing or monotone decreasing. An array is monotone increasing if all its elements from left to right are non-decreasing. An array is monotone decreasing if all  its elements from left to right are non-increasing. Given an integer array return true if the given array is monotonic, or false otherwise.

*/

package main

func isMonotonic(A []int) bool {
	if len(A) < 2 {
		return true
	}

	inc, dec := true, true
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			dec = false
		} else if A[i] < A[i-1] {
			inc = false
		}
	}

	return dec || inc
}
