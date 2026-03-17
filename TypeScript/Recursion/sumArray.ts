// Sum of array elements
//
// No loops
//
// Example: [1, 2, 3, 4] → 10

function sumArray(arr: number[], i: number): number {
    if (i >= arr.length) {
        return 0;
    }
    return arr[i] + sumArray(arr, i + 1);
}