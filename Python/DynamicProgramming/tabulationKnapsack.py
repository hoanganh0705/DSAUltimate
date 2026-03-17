def knapSack(W, wt, val, n):
    # dp[i][j] = max value using first i items with capacity j
    dp = [[0] * (W + 1) for _ in range(n + 1)]

    for i in range(1, n + 1):
        for j in range(1, W + 1):
            # exclude current item
            exclude = dp[i - 1][j]

            # include current item (if possible)
            include = 0
            if wt[i - 1] <= j:
                include = val[i - 1] + dp[i - 1][j - wt[i - 1]]

            dp[i][j] = max(exclude, include)

    return dp[n][W]