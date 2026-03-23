package main

func leastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	maxFreq := 0
	numberMaxFreq := 0

	for _, task := range tasks {
		index := task - 'A'
		count[index]++

		if maxFreq < count[index] {
			maxFreq = count[index]
			numberMaxFreq = 1
		} else if maxFreq == count[index] {
			numberMaxFreq++
		}
	}

	parts := maxFreq - 1
	slotsPerPart := n - (numberMaxFreq - 1)
	totalSlots := parts * slotsPerPart
	remainingTasks := len(tasks) - maxFreq*numberMaxFreq

	idles := max(0, totalSlots-remainingTasks)

	return len(tasks) + idles
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
