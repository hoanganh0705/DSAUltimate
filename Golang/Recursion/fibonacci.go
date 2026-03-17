// Fibonacci Sequence Implementation in Go (Recursive)
//
// The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones.
// It starts with 0 and 1.
//
// Example:
// Input: n = 5
// Output: 0, 1, 1, 2, 3
//
// Time Complexity: O(2^n) (inefficient due to repeated calculations)
// Space Complexity: O(n) (due to the recursion stack)
//

package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func printFibonacci(n int, current int, next int) {
	if n == 0 {
		return
	}

	fmt.Print(current)

	if n > 1 {
		fmt.Print(", ")
	}

	printFibonacci(n-1, next, current+next)
}
