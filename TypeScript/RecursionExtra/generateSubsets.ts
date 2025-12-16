// Generate All Subsets of a Set in TypeScript
//
// Generating all subsets of a set involves creating the power set, which is the set of all subsets.
// This can be done recursively by branching at each step: include or exclude the current element.
//
// Example:
// Input: [1, 2, 3]
// Output: [[], [1], [2], [3], [1,2], [1,3], [2,3], [1,2,3]]
//
// Time Complexity: O(2^n)
// Space Complexity: O(n) (due to the recursion stack)
//
// Note: Each recursive call branches into two, creating a call tree.

const generateSubsets = (arr: number[]): number[][] => {
    if (arr.length === 0) {
        return [[]];
    }

    const firstElement = arr[0];
    const restOfArray = arr.slice(1);

    const subsetsWithoutFirst = generateSubsets(restOfArray);
    const subsetsWithFirst = subsetsWithoutFirst.map(subset => [firstElement, ...subset]);

    return [...subsetsWithoutFirst, ...subsetsWithFirst]; 
};

// Example usage:
const inputArray = [1, 2, 3];
const subsets = generateSubsets(inputArray);
console.log(subsets);