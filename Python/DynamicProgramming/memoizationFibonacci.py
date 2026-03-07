def memoizationFibonacci(n:int, ht:dict={0:0, 1:1}) -> int:
    if n in ht:
        return ht[n]
    else:
        ht[n] = memoizationFibonacci(n-1, ht) +memoizationFibonacci(n-2, ht)
        return ht[n]
    