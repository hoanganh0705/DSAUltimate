/*
Coding Exercise: Container with most water
Question

Container with most Water - You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]). Find two lines that together with the x-axis form a container, such that the container contains the most water(depth is constant across containers). Return the area(that the 2 lines and the X axis make) of container which can store the max amount of water. Notice that you may not slant the container.

Try

To optimise your solution . In case you approached this question with the Brute force method your Time Complexity might be O(n^2).

Try to write a better solution with Time Complexity O(n)

*/

package main

func maxArea(array []int) int {
	maxArea := 0

	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			height := array[i]
			if array[j] < height {
				height = array[j]
			}

			area := height * (j - i)

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
