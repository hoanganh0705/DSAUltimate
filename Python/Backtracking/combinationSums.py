'''
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different. (You will not be given an empty candidates array)
'''

def combinationSum(candidates, target):
    res = []
    n = len(candidates)

    def helper(start_index, curr, sum_included):
        # base case
        if sum_included > target:
            return

        if sum_included == target:
            res.append(curr[:])  # copy list
            return

        # recursive case
        for j in range(start_index, n):
            curr.append(candidates[j])
            helper(j, curr, sum_included + candidates[j])  # allow reuse
            curr.pop()

    helper(0, [], 0)
    return res