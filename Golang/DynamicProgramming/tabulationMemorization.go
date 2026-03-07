package main

func tabulationFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	prev := 0
	curr := 1

	for i := 1; i < n; i++ {
		temp := curr
		curr = prev + curr
		prev = temp
	}

	return curr
}
