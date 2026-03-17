# Sum of array elements
#
# No loops
#
# Example: [1, 2, 3, 4] → 10

def sum_array(arr, i):
    if i >= len(arr):
        return 0
    return arr[i] + sum_array(arr, i + 1)