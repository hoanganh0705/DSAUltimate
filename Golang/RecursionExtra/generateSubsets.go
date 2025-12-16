// Generate All Subsets of a Set in Go
//
// Generating all subsets of a set involves creating the power set, which is the set of all subsets.
// This can be done recursively by branching at each step: include or exclude the current element.
//
// Example:
// Input: [1, 2, 3]
// Output: [[], [1], [2], [3], [1,2], [1,3], [2,3], [1,2,3]]
//
// Time Complexity: O(2^n)
// Space Complexity: O(n) (due to the recursion stack)
//
// Note: Each recursive call branches into two, creating a call tree.

package main

func generateSubsets(nums []int) [][]int {
	// base case: mảng rỗng có 1 subset là mảng rỗng
	if len(nums) == 0 {
		return [][]int{{}}
	}

	first := nums[0]
	rest := nums[1:]

	// subsets không có phần tử first
	subsetsWithoutFirst := generateSubsets(rest)

	// subsets có phần tử first
	subsetsWithFirst := make([][]int, 0, len(subsetsWithoutFirst))

	for _, subset := range subsetsWithoutFirst {
		// tạo subset mới = first + subset
		newSubset := append([]int{first}, subset...)
		subsetsWithFirst = append(subsetsWithFirst, newSubset)
	}

	// gộp 2 nhóm lại
	return append(subsetsWithoutFirst, subsetsWithFirst...)
}
