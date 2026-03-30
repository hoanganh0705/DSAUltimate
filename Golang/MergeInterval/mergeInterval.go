package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		last := result[len(result)-1]

		if start <= last[1] {
			if end > last[1] {
				last[1] = end
			}
		} else {
			result = append(result, []int{start, end})
		}
	}

	return result
}
