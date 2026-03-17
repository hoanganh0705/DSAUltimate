// Factorial
//
// Classic for a reason
//
// factorial(5) → 120

function factorial(n: number): number {
    if (n <= 0) {
        return 1;
    }
    return factorial(n - 1) * n;
}