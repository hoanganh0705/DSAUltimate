"""
Coding Exercise: Sorted Squared Array
Question

You are given an array of Integers in which each subsequent value is not less than the previous value. Write a function that takes this array as an input and returns a new array with the squares of each number sorted in ascending order.

Remember

You can solve this question in multiple ways.

Think about the following -

1.What would be the brute force way of solving this question ? What would be the Time and Space complexity of this approach?

2.Is there a better way to solve this with a more optimum Time complexity than the Brute force way ? 
"""
import numpy as np

def sortSquare (data):
    n = len(data)
    result = np.zeros(n, dtype=int)
    i=0
    j = n - 1
    k = n - 1
    while i <= j :
        if data[i]**2 > data[j] **2:
            result[k] = data[i]**2
            i +=1
        else:
            result[k]= data[j] **2
            j -=1
        k-=1
    return result
    
data = [-9, -5, 2, 3, 4, 10]
print(sortSquare(data)) 