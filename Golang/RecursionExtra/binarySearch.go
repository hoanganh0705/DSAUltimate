// Binary Search Implementation in Go
//
// Binary search is an efficient algorithm for finding an item from a sorted list of items.
// It works by repeatedly dividing in half the portion of the list that could contain the item,
// until you've narrowed down the possible locations to just one.
//
// Time Complexity: O(log n)
// Space Complexity: O(1)
//
// Example:
// Input: sorted array = [1, 3, 5, 7, 9], target = 5
// Output: 2 (index of 5)
//
// Input: sorted array = [1, 3, 5, 7, 9], target = 6
// Output: -1 (not found)

package main

func binarySearch(arr *[]int, target int, left int, right int) int {
	// base case: không còn khoảng tìm kiếm
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if (*arr)[mid] == target {
		return mid
	}

	if target < (*arr)[mid] {
		return binarySearch(arr, target, left, mid-1)
	}

	return binarySearch(arr, target, mid+1, right)
}
