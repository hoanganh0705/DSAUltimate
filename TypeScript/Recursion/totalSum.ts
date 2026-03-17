/*

Bài tập đệ quy: Tính tổng chữ số của một số nguyên dương

Đề bài:
Viết hàm đệ quy sumDigits(n) trả về tổng các chữ số của số nguyên dương n.

Ví dụ:

sumDigits(1234) = 1 + 2 + 3 + 4 = 10

sumDigits(505) = 5 + 0 + 5 = 10

*/

const sumDigits = (n: number): number => {
    if (n === 0) {
        return 0;
    } else {
        return (n % 10) + sumDigits(Math.floor(n / 10));
    }
};

// Test cases
console.log(sumDigits(1234)); // Output: 10
console.log(sumDigits(505));  // Output: 10
console.log(sumDigits(0));    // Output: 0
console.log(sumDigits(9));    // Output: 9
console.log(sumDigits(1001)); // Output: 2