'''

Coding Exercise: Monotonic Array
Question

An array is monotonic if it is either monotone increasing or monotone decreasing. An array is monotone increasing if all its elements from left to right are non-decreasing. An array is monotone decreasing if all  its elements from left to right are non-increasing. Given an integer array return true if the given array is monotonic, or false otherwise.

'''

def isMonotonic(arr):
    if len(arr) < 2:
        return True

    inc = True  # giả sử dãy tăng
    dec = True  # giả sử dãy giảm

    for i in range(len(arr) - 1):
        if arr[i] > arr[i + 1]:
            inc = False
        elif arr[i] < arr[i + 1]:
            dec = False

    return inc or dec


print(isMonotonic([1, 2, 2, 3])) 
print(isMonotonic([6, 5, 4, 4])) 
print(isMonotonic([1, 3, 2]))    
print(isMonotonic([1, 1, 1])) 
