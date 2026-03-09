'''
You are given an integer array cost where cost[i] is the cost of the i-th step on a staircase.
Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.

Example 1:
Input: cost = [10, 20, 30]
Output: 20

Explanation:
You will start at index 1.
- Pay 20 and climb two steps to reach the top.
'''

def minCostClimbingStairs(cost):
    # write code here
    n = len(cost)

    def helper(index):
        # return the min cost for reaching the top starting from step - index

        # base
        if index > n - 1:
            return 0

        # recursive case

        # one step
        onestep = cost[index] + helper(index + 1)

        # two steps
        twostep = cost[index] + helper(index + 2)

        return min(onestep, twostep)

    return min(helper(0), helper(1))