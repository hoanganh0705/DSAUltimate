package main

import "fmt"

func main() {
	// sum digit
	// result := sumDigits(999, 0)
	// fmt.Println(result)

	// subsequences string

	// result := Subsequences("abc")
	// fmt.Printf("%q\n", result)

	// result := sumOfFirstN(5)
	// fmt.Println(result)

	// fmt.Println(factorial(5))

	// arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(sumPtr(&arr, 0))

	n := 5
	result := fibonacci(n)
	fmt.Println(result)

	printFibonacci(n, 0, 1)
}
