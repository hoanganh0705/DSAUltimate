package main

func combine(n int, k int) [][]int {
	var res [][]int
	var helper func(start int, curr []int)

	helper = func(start int, curr []int) {
		if len(curr) == k {
			combination := make([]int, k)
			copy(combination, curr)
			res = append(res, combination)
			return
		}

		for j := start; j <= n; j++ {
			curr = append(curr, j)
			helper(j+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	helper(1, []int{})
	return res
}
