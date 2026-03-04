/*


Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different. (You will not be given an empty candidates array)
*/

function combinationSum(candidates: number[], target: number): number[][] {
    const res: number[][] = [];
    const n = candidates.length;

    function helper(startIndex: number, curr: number[], sumIncluded: number) {
        // base case
        if (sumIncluded > target) {
            return;
        }

        if (sumIncluded === target) {
            res.push([...curr]); // clone array
            return;
        }

        // recursive case
        for (let j = startIndex; j < n; j++) {
            curr.push(candidates[j]);
            helper(j, curr, sumIncluded + candidates[j]); // reuse element
            curr.pop();
        }
    }

    helper(0, [], 0);
    return res;
}