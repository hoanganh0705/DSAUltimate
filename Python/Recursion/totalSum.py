'''
Bài tập đệ quy: Tính tổng chữ số của một số nguyên dương

Đề bài:
Viết hàm đệ quy sumDigits(n) trả về tổng các chữ số của số nguyên dương n.

Ví dụ:

sumDigits(1234) = 1 + 2 + 3 + 4 = 10

sumDigits(505) = 5 + 0 + 5 = 10


'''

def sumDigits(n):
    if n < 10:
        return n
    return sumDigits(n // 10) + (n % 10)

print(sumDigits(1234)) 
