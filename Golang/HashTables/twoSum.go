/*
Coding Exercise: Two Sum

Question:

Two Sum - You are given an array of Integers and another integer targetValue. Write a function that will take these inputs and return the indices of the 2 integers in the array that add up targetValue.

Try:

Try to optimise your solution and arrive at a Time Complexity of O(n)
*/

package main

func twoSum(array []int, targetValue int) (int, int) {
	nums := make(map[int]int)
	for i, num := range array {
		complement := targetValue - num
		if j, ok := nums[complement]; ok {
			return j, i
		}
		nums[num] = i
	}
	return -1, -1
}
