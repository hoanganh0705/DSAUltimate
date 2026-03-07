package main

func memoizationFibonacci(n int, ht map[int]int) int {
	if val, ok := ht[n]; ok {
		return val
	}

	ht[n] = memoizationFibonacci(n-1, ht) + memoizationFibonacci(n-2, ht)
	return ht[n]
}
