package main

import (
	"sort"
	"strconv"
	"strings"
)

func threeSumNew(nums []int) [][]int {
	res := make(map[string]bool)

	for i := 0; i < len(nums); i++ {
		need := make(map[int]bool)

		for j := i + 1; j < len(nums); j++ {
			valueNeeded := -(nums[i] + nums[j])

			if need[valueNeeded] {
				triplet := []int{nums[i], nums[j], valueNeeded}
				sort.Ints(triplet)

				key := strconv.Itoa(triplet[0]) + "," +
					strconv.Itoa(triplet[1]) + "," +
					strconv.Itoa(triplet[2])

				res[key] = true
			}

			need[nums[j]] = true
		}
	}

	result := [][]int{}
	for key := range res {
		parts := strings.Split(key, ",")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])

		result = append(result, []int{a, b, c})
	}

	return result
}
