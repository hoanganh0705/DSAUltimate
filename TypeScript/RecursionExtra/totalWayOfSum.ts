/*

Problem:
Given an array of positive integers nums and a target T,
return the number of ways to pick elements (each element can be used multiple times) so they sum to exactly T.

Example:
nums = [1, 3, 4], T = 5
Possible ways:

1+1+1+1+1

1+1+3

1+4

3+1+1
(etc… order matters or doesn’t matter—you pick, but be consistent)

*/

function totalWayOfSum(nums: number[], target: number): number {
    // Base case: if target is 0, there is one way to make the sum (by choosing nothing)
    if (target === 0) {
        return 1;
    }

    // If target is negative, no way to make the sum
    if (target < 0) {
        return 0;
    }

    let totalWays = 0;

    // Iterate through each number in the array
    for (let num of nums) {
        // Recursively call the function with reduced target
        totalWays += totalWayOfSum(nums, target - num);
    }

    return totalWays;
}