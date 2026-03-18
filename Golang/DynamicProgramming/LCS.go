package main

func longestCommonSubsequence(text1 string, text2 string) int {

	n := len(text1)
	m := len(text2)

	var helper func(index1 int, index2 int) int
	helper = func(index1 int, index2 int) int {
		// base case
		if index1 >= n || index2 >= m {
			return 0
		}

		// match
		if text1[index1] == text2[index2] {
			return 1 + helper(index1+1, index2+1)
		}

		// no match
		return max(
			helper(index1, index2+1),
			helper(index1+1, index2),
		)
	}

	return helper(0, 0)
}
