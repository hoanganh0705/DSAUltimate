'''

Coding Exercise (k-th symbol in Grammar)
We build a table of n rows (1-indexed). We start by writing 0 in the 1st row. Now in every subsequent row, we look at the previous row and replace each 
occurrence of 0 with 01, and each occurrence of 1 with 10. For example, for n = 3, the 1st row is 0, the 2nd row is 01, and the 3rd row is 0110. Given two 
integer n and k, return the kth (1-indexed) symbol in the nth row of a table of n rows.

'''

def kthGrammar(n: int, k: int) -> int:
    if n == 1:
        return 0

    length = 2 ** (n - 1)      # độ dài hàng n
    half = length // 2         # nửa đầu

    if k <= half:
        # nửa đầu: giống y hàng n-1
        return kthGrammar(n - 1, k)
    else:
        # nửa sau: đảo bit của kết quả tương ứng ở hàng n-1
        return 1 - kthGrammar(n - 1, k - half)


# ví dụ đúng:
print(kthGrammar(3, 1))  # 0
print(kthGrammar(3, 2))  # 1
print(kthGrammar(3, 3))  # 1
print(kthGrammar(3, 4))  # 0
