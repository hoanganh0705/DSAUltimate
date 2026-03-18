def memoizedLCS(text1, text2):
    n = len(text1)
    m = len(text2)

    # dp[i][j] = LCS from text1[i:], text2[j:]
    dp = [[-1] * m for _ in range(n)]

    def helper(i, j):
        # base case
        if i >= n or j >= m:
            return 0

        if dp[i][j] != -1:
            return dp[i][j]

        # match
        if text1[i] == text2[j]:
            dp[i][j] = 1 + helper(i + 1, j + 1)
        else:
            # no match
            dp[i][j] = max(
                helper(i, j + 1),
                helper(i + 1, j)
            )

        return dp[i][j]

    return helper(0, 0)