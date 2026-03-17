/*

Bài tập đệ quy: Tính tổng chữ số của một số nguyên dương

Đề bài:
Viết hàm đệ quy sumDigits(n) trả về tổng các chữ số của số nguyên dương n.

Ví dụ:

sumDigits(1234) = 1 + 2 + 3 + 4 = 10

sumDigits(505) = 5 + 0 + 5 = 10

*/

// this is the way that I use my brain to think but it kinda sucks, because recursion usually avoid accumulator
package main

func sumDigits(number int, counter int) int {
	// base case
	if number/10 == 0 {
		return number%10 + counter
	} else {
		counter += number % 10
		return sumDigits(number/10, counter)
	}

}

func sumDigitsBestWay(number int) int {
	if number < 10 {
		return number
	}

	return sumDigitsBestWay(number/10) + (number % 10)
}
