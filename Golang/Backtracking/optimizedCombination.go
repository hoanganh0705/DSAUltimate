package main

func optimizedCombine(n int, k int) [][]int {
	var res [][]int

	var helper func(start int, curr []int)
	helper = func(start int, curr []int) {
		if len(curr) == k {
			comb := make([]int, k)
			copy(comb, curr)
			res = append(res, comb)
			return
		}

		need := k - len(curr)

		for j := start; j <= n-(need-1); j++ {
			curr = append(curr, j)
			helper(j+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	helper(1, []int{})
	return res
}
