package main

import "fmt"

func main() {
	n := 3
	moves := toh(n, "rod 1", "rod 3", "rod 2")
	fmt.Println(moves)

	nestedArr := []interface{}{1, 2, []interface{}{7, []interface{}{3, 4}, 2}}
	result := powersum(nestedArr, 0)
	fmt.Println(result)
}
