# Reverse String Implementation in Python
#
# Reversing a string means rearranging its characters in the opposite order.
# For example, "hello" becomes "olleh".
#
# In Python, there are multiple ways to reverse a string:
# 1. Using slicing: s[::-1]
# 2. Using reversed() function: ''.join(reversed(s))
# 3. Using a loop to build the reversed string
# 4. Using two pointers to swap in a list
#
# Time Complexity: O(n)
# Space Complexity: O(n) due to the list
#
# Example:
# Input: "hello"
# Output: "olleh"
#
# Input: "world"
# Output: "dlrow"
#
# Input: "a"
# Output: "a" (single character remains the same)
#
# Input: ""
# Output: "" (empty string remains empty)