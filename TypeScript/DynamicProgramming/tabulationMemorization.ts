const tabulationFibonacci = (n: number): number => {
    if (n <= 1) return n
    let prev = 0
    let curr = 1
    let counter = 1
    while (counter < n) {
        const temp = curr
        curr = prev + curr
        prev = temp
        counter++
    }
    return curr
}

console.log(tabulationFibonacci(10))