package main

func reverse(array []int, start int, end int) {
	for start < end {
		array[start], array[end] = array[end], array[start]
		start++
		end--
	}
}

func enhancedRotateArray(array []int, k int) []int {
	if len(array) == 0 {
		return []int{}
	}

	k = k % len(array)
	if k == 0 {
		return array
	}

	reverse(array, 0, len(array)-1)
	reverse(array, 0, k-1)
	reverse(array, k, len(array)-1)

	return array
}
