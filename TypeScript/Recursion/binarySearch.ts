// Binary Search Implementation in TypeScript
//
// Binary search is an efficient algorithm for finding an item from a sorted list of items.
// It works by repeatedly dividing in half the portion of the list that could contain the item,
// until you've narrowed down the possible locations to just one.
//
// Time Complexity: O(log n)
// Space Complexity: O(1)
//
// Example:
// Input: sorted array = [1, 3, 5, 7, 9], target = 5
// Output: 2 (index of 5)
//
// Input: sorted array = [1, 3, 5, 7, 9], target = 6
// Output: -1 (not found)

function binarySearch(arr: number[], target: number): number {
    let left: number = 0;
    let right: number = arr.length - 1;

    while (left <= right) {
        const mid: number = Math.floor((left + right) / 2);

        if (arr[mid] === target) {
            return mid; // Target found
        } else if (arr[mid] < target) {
            left = mid + 1; // Search in the right half
        } else {
            right = mid - 1; // Search in the left half
        }
    }

    return -1; // Target not found
}

// Example usage:
console.log(binarySearch([1, 3, 5, 7, 9], 5)); // Output: 2
console.log(binarySearch([1, 3, 5, 7, 9], 6)); // Output: -1