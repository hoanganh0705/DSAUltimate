/*

Frequency of the Most Frequent Element
The frequency of an element is the number of times it occurs in an array.

You are given an integer array nums and an integer k. In one operation, you can choose an index of nums and increment the element at that index by 1.

Return the maximum possible frequency of an element after performing at most k operations.

Example:

Input: nums = [2,3,5], k = 5

Output: 3

Explanation: Increment 2 to 5 (3 operations), then increment 3 to 5 (2 operations). The array becomes [5,5,5].
*/
function maxFrequency(nums: number[], k: number): number {
    nums.sort((a, b) => a - b)

    let left = 0
    let total = 0
    let res = 0

    for (let right = 0; right < nums.length; right++) {
        total += nums[right]

        while (nums[right] * (right - left + 1) > total + k) {
            total -= nums[left]
            left++
        }

        res = Math.max(res, right - left + 1)
    }

    return res
}