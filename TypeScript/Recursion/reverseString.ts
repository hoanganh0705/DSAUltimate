// Reverse String Implementation in TypeScript
//
// Reversing a string means rearranging its characters in the opposite order.
// For example, "hello" becomes "olleh".
//
// In TypeScript/JavaScript, there are multiple ways to reverse a string:
// 1. Using built-in methods: s.split('').reverse().join('')
// 2. Using a loop to build the reversed string
// 3. Using two pointers to swap in an array
//
// Time Complexity: O(n)
// Space Complexity: O(n) due to the array
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

function reverseString(s: string): string {
    // Convert the string to an array of characters
    const charArray: string[] = s.split('');
    
    // Initialize two pointers
    let left: number = 0;
    let right: number = charArray.length - 1;
    
    // Swap characters until the pointers meet in the middle
    while (left < right) {
        // Swap characters at left and right pointers
        const temp: string = charArray[left];
        charArray[left] = charArray[right];
        charArray[right] = temp;
        
        // Move the pointers closer to the center
        left++;
        right--;
    }
    
    // Convert the array back to a string and return
    return charArray.join('');
}

// Example usage:
console.log(reverseString("hello")); // Output: "olleh"
console.log(reverseString("world")); // Output: "dlrow"
console.log(reverseString("a"));     // Output: "a"
console.log(reverseString(""));      // Output: ""