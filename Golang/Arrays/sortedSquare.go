/*
Coding Exercise: Sorted Squared Array
Question

You are given an array of Integers in which each subsequent value is not less than the previous value. Write a function that takes this array as an input and returns a new array with the squares of each number sorted in ascending order.

Remember

You can solve this question in multiple ways.

Think about the following -

1.What would be the brute force way of solving this question ? What would be the Time and Space complexity of this approach?

2.Is there a better way to solve this with a more optimum Time complexity than the Brute force way ?

*/

package main

func sortedSquare(arr []int) []int {
	n := len(arr)
	result := make([]int, n)

	i := 0
	j := n - 1
	k := n - 1

	for i <= j {
		leftSq := arr[i] * arr[i]
		rightSq := arr[j] * arr[j]

		if leftSq > rightSq {
			result[k] = leftSq
			i++
		} else {
			result[k] = rightSq
			j--
		}
		k--
	}

	return result
}
