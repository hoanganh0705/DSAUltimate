// Find max in array
//
// Recursively compare elements
//
// No sorting, no loops

function findMax(arr: number[], i: number): number {
    if (arr.length === 0) {
        throw new Error("findMax: empty array");
    }
    // base case: last element
    if (i === arr.length - 1) {
        return arr[i];
    }
    // max of the rest
    const m = findMax(arr, i + 1);
    if (arr[i] > m) {
        return arr[i];
    }
    return m;
}

function max(arr: number[]): number {
    return findMax(arr, 0);
}