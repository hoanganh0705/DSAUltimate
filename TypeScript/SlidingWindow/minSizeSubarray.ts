/*
Min Size Subarray Sum
Minimum Size Subarray Sum:

Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example:

target = 15

nums = [5, 1, 3, 5, 10, 7, 4, 9, 2, 8]

expected = 2
*/

function minSubArrayLen(target: number, nums: number[]): number {
    let left = 0;
    let currentSum = 0;
    let minLen = Infinity;

    for (let right = 0; right < nums.length; right++) {
        currentSum += nums[right];

        while (currentSum >= target) {
            minLen = Math.min(minLen, right - left + 1);
            currentSum -= nums[left];
            left++;
        }
    }

    return minLen !== Infinity ? minLen : 0;
}