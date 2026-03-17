# Find max in array
#
# Recursively compare elements
#
# No sorting, no loops

def find_max(arr, i):
    if len(arr) == 0:
        raise ValueError("find_max: empty list")
    # base case: last element
    if i == len(arr) - 1:
        return arr[i]
    # max of the rest
    m = find_max(arr, i + 1)
    if arr[i] > m:
        return arr[i]
    return m

def max(arr):
    return find_max(arr, 0)