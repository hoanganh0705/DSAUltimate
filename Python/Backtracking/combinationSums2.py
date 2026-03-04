'''
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

'''

def combinationSum2(candidates, target):
    candidates.sort()
    res = []
    n = len(candidates)

    def helper(index, curr, curr_sum):
        # base case
        if curr_sum == target:
            res.append(curr[:])
            return

        if curr_sum > target:
            return

        if index > n - 1:
            return

        # recursive case
        used = {}

        for i in range(index, n):
            if candidates[i] not in used:
                used[candidates[i]] = True

                curr.append(candidates[i])
                helper(i + 1, curr, curr_sum + candidates[i])
                curr.pop()

    helper(0, [], 0)
    return res