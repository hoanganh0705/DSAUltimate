def longestCommonSubsequence(text1, text2):
    n = len(text1)
    m = len(text2)

    def helper(index1, index2):
        # base case
        if index1 >= n or index2 >= m:
            return 0

        # match
        if text1[index1] == text2[index2]:
            return 1 + helper(index1 + 1, index2 + 1)

        # no match
        return max(
            helper(index1, index2 + 1),
            helper(index1 + 1, index2)
        )

    return helper(0, 0)