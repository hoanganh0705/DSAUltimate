package main

import "fmt"

func main() {
	arr := [][]int{{70, 10}, {90, 20}, {150, 30}}
	W := 25
	n := len(arr)
	fmt.Println(fractionalKnapsack(W, arr, n))
}
