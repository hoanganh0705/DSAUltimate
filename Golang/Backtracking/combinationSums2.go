/*
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

*/

package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var res = [][]int{}
	var n = len(candidates)

	var helper func(index int, curr []int, currSum int)
	helper = func(index int, curr []int, currSum int) {
		// base case
		if currSum == target {
			res = append(res, append([]int(nil), curr...))
			return
		}

		if currSum > target {
			return
		}

		if index > n-1 {
			return
		}

		// recursive case
		used := make(map[int]bool)

		for i := index; i < n; i++ {
			if !used[candidates[i]] {
				used[candidates[i]] = true

				curr = append(curr, candidates[i])
				helper(i+1, curr, currSum+candidates[i])
				curr = curr[:len(curr)-1]
			}
		}
	}

	helper(0, []int{}, 0)
	return res
}
