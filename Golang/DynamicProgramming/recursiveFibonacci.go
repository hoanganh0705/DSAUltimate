// bài toán climbing chairs cũng giống như bài toán fibonacci

package main

func recursiveFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
}
