def memoizedLIS(nums):
    n = len(nums)

    # dp[curr][prev+1] → shift prev by +1 to handle -1
    dp = [[-1] * (n + 1) for _ in range(n)]

    def helper(curr, prev):
        if curr >= n:
            return 0

        if dp[curr][prev + 1] != -1:
            return dp[curr][prev + 1]

        # exclude
        exclude = helper(curr + 1, prev)

        # include
        include = 0
        if prev == -1 or nums[curr] > nums[prev]:
            include = 1 + helper(curr + 1, curr)

        dp[curr][prev + 1] = max(exclude, include)
        return dp[curr][prev + 1]

    return helper(0, -1)