// Reverse String Implementation in Go
//
// Reversing a string means rearranging its characters in the opposite order.
// For example, "hello" becomes "olleh".
//
// There are multiple ways to reverse a string:
// 1. Using built-in functions (if available)
// 2. Using two pointers to swap characters in place
// 3. Using recursion
// 4. Converting to rune slice and reversing
//
// Time Complexity: O(n)
// Space Complexity: O(n) due to rune slice
//
// Example:
// Input: "hello"
// Output: "olleh"
//
// Input: "world"
// Output: "dlrow"
//
// Input: "a"
// Output: "a" (single character remains the same)
//
// Input: ""
// Output: "" (empty string remains empty)

package main
