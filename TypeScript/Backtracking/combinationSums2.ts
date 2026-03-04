/*
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

*/

function combinationSum2(candidates: number[], target: number): number[][] {
    candidates.sort((a, b) => a - b);

    const res: number[][] = [];
    const n = candidates.length;

    function helper(index: number, curr: number[], currSum: number) {

        // base case
        if (currSum === target) {
            res.push([...curr]);
            return;
        }

        if (currSum > target) {
            return;
        }

        if (index > n - 1) {
            return;
        }

        // recursive case
        const used = new Set<number>();

        for (let i = index; i < n; i++) {
            if (!used.has(candidates[i])) {
                used.add(candidates[i]);

                curr.push(candidates[i]);
                helper(i + 1, curr, currSum + candidates[i]);
                curr.pop();
            }
        }
    }

    helper(0, [], 0);
    return res;
}