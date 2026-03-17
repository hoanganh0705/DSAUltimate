// Fibonacci Sequence Implementation in TypeScript (Recursive)
//
// The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones.
// It starts with 0 and 1.
//
// Example:
// Input: n = 5
// Output: 0, 1, 1, 2, 3
//
// Time Complexity: O(2^n) (inefficient due to repeated calculations)
// Space Complexity: O(n) (due to the recursion stack)
//
// Note: This implementation is inefficient and is meant to demonstrate recursion.
// Optimization can be achieved using memoization or dynamic programming.

const fibonacci = (n: number): number => {
    if (n <= 0) return 0;
    if (n === 1) return 1;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

const printFibonacci = (n: number, a: number = 0, b: number = 1): void => {
    if (n <= 0) return;
    console.log(a);
    printFibonacci(n - 1, b, a + b);
}