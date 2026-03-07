def tabulationFibonacci(n: int) -> int:
    if n <=1: return n
    prev = 0 
    curr = 1
    counter = 1
    while counter < n:
        temp = curr
        curr = prev + curr
        prev = temp
        counter += 1
    return curr

print(tabulationFibonacci(10))