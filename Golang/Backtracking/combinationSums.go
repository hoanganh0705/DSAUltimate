/*


Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different. (You will not be given an empty candidates array)
*/

package main

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	n := len(candidates)

	var helper func(start int, curr []int, sumIncluded int)

	helper = func(start int, curr []int, sumIncluded int) {
		if sumIncluded > target {
			return
		}

		if sumIncluded == target {
			res = append(res, append([]int(nil), curr...))
			return
		}

		for j := start; j < n; j++ {
			curr = append(curr, candidates[j])
			helper(j, curr, sumIncluded+candidates[j]) // reuse element
			curr = curr[:len(curr)-1]
		}
	}

	helper(0, []int{}, 0)
	return res
}
