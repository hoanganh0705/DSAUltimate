/*

Frequency of the Most Frequent Element
The frequency of an element is the number of times it occurs in an array.

You are given an integer array nums and an integer k. In one operation, you can choose an index of nums and increment the element at that index by 1.

Return the maximum possible frequency of an element after performing at most k operations.

Example:

Input: nums = [2,3,5], k = 5

Output: 3

Explanation: Increment 2 to 5 (3 operations), then increment 3 to 5 (2 operations). The array becomes [5,5,5].
*/

package main

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	left := 0
	total := 0
	res := 0

	for right := range nums {
		total += nums[right]

		for nums[right]*(right-left+1) > total+k {
			total -= nums[left]
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}

	return res
}
