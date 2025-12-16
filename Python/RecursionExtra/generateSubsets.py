# Generate All Subsets of a Set in Python
#
# Generating all subsets of a set involves creating the power set, which is the set of all subsets.
# This can be done recursively by branching at each step: include or exclude the current element.
#
# Example:
# Input: [1, 2, 3]
# Output: [[], [1], [2], [3], [1,2], [1,3], [2,3], [1,2,3]]
#
# Time Complexity: O(2^n)
# Space Complexity: O(n) (due to the recursion stack)
#
# Note: Each recursive call branches into two, creating a call tree.

def generateSubsets(arr):
    if len(arr) == 0:
        return [[]] 
    
    first = arr[0]
    rest_subsets = generateSubsets(arr[1:])
    subsets_with_first = [[first] + subset for subset in rest_subsets]
    return rest_subsets + subsets_with_first

# Example usage:
input_set = [1, 2, 3]
subsets = generateSubsets(input_set)
print("All subsets:", subsets)