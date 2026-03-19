def tabulationLIS(nums):
    n = len(nums)

    # dp[i][j] → j represents prev_index + 1
    dp = [[0] * (n + 1) for _ in range(n + 1)]

    for i in range(n - 1, -1, -1):
        for j in range(i, -1, -1):
            # exclude
            exclude = dp[i + 1][j]

            # include
            include = 0
            if j == 0 or nums[i] > nums[j - 1]:
                include = 1 + dp[i + 1][i + 1]

            dp[i][j] = max(exclude, include)

    return dp[0][0]