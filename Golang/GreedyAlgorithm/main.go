package main

import "fmt"

func main() {
	arr := [][]int{{70, 10}, {90, 20}, {150, 30}}
	W := 25
	n := len(arr)
	fmt.Println(fractionalKnapsack(W, arr, n))

	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n1 := 2
	fmt.Println(leastInterval(tasks, n1))
}
