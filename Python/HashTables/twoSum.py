'''

Coding Exercise: Two Sum

Question:

Two Sum - You are given an array of Integers and another integer targetValue. Write a function that will take these inputs and return the indices of the 2 integers in the array that add up targetValue.

Try:

Try to optimise your solution and arrive at a Time Complexity of O(n)

'''

def twoSum(array, targetValue):
    nums = {}
    for i, num in enumerate(array):
        complement = targetValue - num
        if complement in nums:
            return (nums[complement], i)
        nums[num] = i
    return (-1, -1)