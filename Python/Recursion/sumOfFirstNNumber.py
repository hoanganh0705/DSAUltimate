def sumOfFirstN(n):
    if n <= 0:
        return 0
    return sumOfFirstN(n - 1) + n