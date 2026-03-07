const memo = new Map<number, number>([
    [0,0],
    [1,1]
])

function fib(n: number): number {
    if (memo.has(n)) return memo.get(n)!

    memo.set(n, fib(n-1) + fib(n-2))
    return memo.get(n)!
}

console.log(fib(10))