/*
Sliding Window Maximum
Sliding Window Maximum

You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.


Example:

Input: nums = [2,3,-2,-4,5,2,8,11], k = 3

Output: [3,3,5,5,8,11]
*/

function maxSlidingWindow(nums: number[], k: number): number[] {
    const dq: number[] = []; // lưu index
    const output: number[] = [];

    for (let i = 0; i < k; i++) {
        while (dq.length && nums[i] >= nums[dq[dq.length - 1]]) {
            dq.pop();
        }
        dq.push(i);
    }

    output.push(nums[dq[0]]);

    for (let i = k; i < nums.length; i++) {
        if (dq.length && dq[0] <= i - k) {
            dq.shift();
        }

        while (dq.length && nums[i] >= nums[dq[dq.length - 1]]) {
            dq.pop();
        }

        dq.push(i);
        output.push(nums[dq[0]]);
    }

    return output;
}