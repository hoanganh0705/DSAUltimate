package main

func maxAreaOptimum(array []int) int {
	left := 0
	right := len(array) - 1
	maxArea := 0

	for left < right {
		height := array[left]
		if array[right] < height {
			height = array[right]
		}

		area := height * (right - left)

		if area > maxArea {
			maxArea = area
		}

		if array[left] < array[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
