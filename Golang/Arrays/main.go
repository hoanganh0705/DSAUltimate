package main

import "fmt"

func main() {
	// sortSquare
	// arr := []int{-9, -5, 2, 3, 4, 10}
	// result := sortedSquare(arr)
	// for i := 0; i < len(result); i++ {
	// 	fmt.Printf("%d ", result[i])
	// }

	monoArr := []int{2, 3, 6, 5, 6}
	monoArr1 := []int{1, 3, 2}
	monoArr2 := []int{1, 1, 1}
	fmt.Println(isMonotonic(monoArr))
	fmt.Println(isMonotonic(monoArr1))
	fmt.Println(isMonotonic(monoArr2))
}
