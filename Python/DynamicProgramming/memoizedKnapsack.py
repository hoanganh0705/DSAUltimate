def memoizedKnapsack(W, wt, val, n):
    # dp table initialized with -1
    dp = [[-1] * (W + 1) for _ in range(n)]

    def helper(index, rem_weight):
        # base case
        if index >= n or rem_weight == 0:
            return 0

        # return cached result
        if dp[index][rem_weight] != -1:
            return dp[index][rem_weight]

        # recursive case
        exclude = helper(index + 1, rem_weight)
        include = 0

        if wt[index] <= rem_weight:
            include = val[index] + helper(index + 1, rem_weight - wt[index])

        # store result
        dp[index][rem_weight] = max(exclude, include)
        return dp[index][rem_weight]

    return helper(0, W)