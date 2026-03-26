package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			left := i + 1
			right := len(nums) - 1

			for left < right {
				sumThree := nums[i] + nums[left] + nums[right]

				if sumThree == 0 {
					res = append(res, []int{nums[i], nums[left], nums[right]})
					left++

					for left < right && nums[left] == nums[left-1] {
						left++
					}

				} else if sumThree < 0 {
					left++
				} else {
					right--
				}
			}
		}
	}

	return res
}
