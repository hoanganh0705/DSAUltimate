/*

Coding Exercise (k-th symbol in Grammar)
We build a table of n rows (1-indexed). We start by writing 0 in the 1st row. Now in every subsequent row, we look at the previous row and replace each
occurrence of 0 with 01, and each occurrence of 1 with 10. For example, for n = 3, the 1st row is 0, the 2nd row is 01, and the 3rd row is 0110. Given two
integer n and k, return the kth (1-indexed) symbol in the nth row of a table of n rows.

*/

// the first half of the row is the same as the previous row and the second half is the inverted version of the first half

package main

func KthSymbol(n, k int) int {
	if n == 1 {
		return 0
	}

	// Độ dài hàng n là 2^(n-1), nửa đầu là 2^(n-2)
	// a << x = a * 2^x
	half := 1 << (n - 2) // 2^(n-2) ) phép dịch bit trái

	if k <= half {
		// Nửa đầu: giống y hàng trước
		return KthSymbol(n-1, k)
	}

	// Nửa sau: là đảo bit của hàng trước tại vị trí (k - half)
	val := KthSymbol(n-1, k-half)
	if val == 0 {
		return 1
	}
	return 0

	// hoặc
	// return 1 - KthSymbol(n-1, k-half)

}
