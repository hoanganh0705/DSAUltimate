# Binary Search Implementation in Python
#
# Binary search is an efficient algorithm for finding an item from a sorted list of items.
# It works by repeatedly dividing in half the portion of the list that could contain the item,
# until you've narrowed down the possible locations to just one.
#
# Time Complexity: O(log n)
# Space Complexity: O(1)
#
# Example:
# Input: sorted array = [1, 3, 5, 7, 9], target = 5
# Output: 2 (index of 5)
#
# Input: sorted array = [1, 3, 5, 7, 9], target = 6
# Output: -1 (not found)

def binary_search(arr: list[int], target: int) -> int:
    left, right = 0, len(arr) - 1
    
    while left <= right:
        mid = left + (right - left) // 2
        
        # Check if target is present at mid
        if arr[mid] == target:
            return mid
        # If target is greater, ignore left half
        elif arr[mid] < target:
            left = mid + 1
        # If target is smaller, ignore right half
        else:
            right = mid - 1
    
    # Target was not found in the array
    return -1