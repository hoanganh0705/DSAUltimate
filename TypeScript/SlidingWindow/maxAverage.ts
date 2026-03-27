/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10^-5 will be accepted.



Example:

Input: [1,13,-6,-3,40,2], k = 4

Output: 11

13 + -6 + -3 + 40 = 44; 44/4 = 11
*/

function findMaxAverage(nums: number[], k: number): number {
    let currSum = 0;

    for (let i = 0; i < k; i++) {
        currSum += nums[i];
    }

    let maxSum = currSum;

    for (let i = k; i < nums.length; i++) {
        currSum = currSum + nums[i] - nums[i - k];
        maxSum = Math.max(maxSum, currSum);
    }

    return maxSum / k;
}