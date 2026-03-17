/*

Sum of first n numbers

Input: n

Output: 1 + 2 + ... + n

Example: sum(5) → 15
*/

package main

func sumOfFirstN(n int) int {
	if n <= 0 {
		return 0
	}

	return sumOfFirstN(n-1) + n
}
