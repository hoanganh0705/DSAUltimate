# Fibonacci Sequence Implementation in Python (Recursive)
#
# The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones.
# It starts with 0 and 1.
#
# Example:
# Input: n = 5
# Output: 0, 1, 1, 2, 3
#
# Time Complexity: O(2^n) (inefficient due to repeated calculations)
# Space Complexity: O(n) (due to the recursion stack)
#
# Note: This implementation is inefficient and is meant to demonstrate recursion.
# Optimization can be achieved using memoization or dynamic programming.

def fibonacci(n, a = 0, b = 1):
    if n == 0:
        return a
    return fibonacci(n - 1, b, a + b)

def printFibonacci(n, a = 0, b = 1):
    if n == 0:
        return
    print(a, end=", " if n > 1 else "\n")
    printFibonacci(n - 1, b, a + b)

# Example usage:
n = 5       
print(fibonacci(n))  # Output: 5
printFibonacci(n)    # Output: 0, 1, 1,