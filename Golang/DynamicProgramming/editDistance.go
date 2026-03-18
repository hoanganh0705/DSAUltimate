package main

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	var helper func(int, int) int
	helper = func(i int, j int) int {
		// base cases
		if i >= n && j >= m {
			return 0
		}
		if i >= n {
			return m - j
		}
		if j >= m {
			return n - i
		}

		if word1[i] == word2[j] {
			return helper(i+1, j+1)
		}

		replace := 1 + helper(i+1, j+1)
		delete := 1 + helper(i+1, j)
		insert := 1 + helper(i, j+1)

		// min of three
		if replace < delete && replace < insert {
			return replace
		} else if delete < insert {
			return delete
		}
		return insert
	}

	return helper(0, 0)
}
