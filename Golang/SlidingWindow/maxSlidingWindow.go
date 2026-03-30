/*
Sliding Window Maximum

You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.


Example:

Input: nums = [2,3,-2,-4,5,2,8,11], k = 3

Output: [3,3,5,5,8,11]
*/

package main

func maxSlidingWindow(nums []int, k int) []int {
	dq := []int{} // lưu index
	result := []int{}

	for i := 0; i < k; i++ {
		for len(dq) > 0 && nums[i] >= nums[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}

	result = append(result, nums[dq[0]])

	for i := k; i < len(nums); i++ {
		if len(dq) > 0 && dq[0] <= i-k {
			dq = dq[1:]
		}

		for len(dq) > 0 && nums[i] >= nums[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i)
		result = append(result, nums[dq[0]])
	}

	return result
}
