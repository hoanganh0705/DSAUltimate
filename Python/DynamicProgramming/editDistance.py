def minDistance(word1, word2):
    n = len(word1)
    m = len(word2)

    def helper(i, j):
        # base cases
        if i >= n and j >= m:
            return 0
        if i >= n:
            return m - j
        if j >= m:
            return n - i

        # if characters match
        if word1[i] == word2[j]:
            return helper(i + 1, j + 1)

        # operations
        replace = 1 + helper(i + 1, j + 1)
        delete = 1 + helper(i + 1, j)
        insert = 1 + helper(i, j + 1)

        return min(replace, delete, insert)

    return helper(0, 0)