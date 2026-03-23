package main

func rotateArray(array []int, k int) []int {
	if len(array) == 0 {
		return []int{}
	}

	k = k % len(array)

	temp := make([]int, k)
	copy(temp, array[len(array)-k:])

	for i := len(array) - k - 1; i >= 0; i-- {
		array[i+k] = array[i]
	}

	copy(array, temp)

	return array
}
