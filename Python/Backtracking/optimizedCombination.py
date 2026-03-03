def optimizedCombine(n, k):
    res = []

    def helper(start, curr):
        if len(curr) == k:
            res.append(curr[:])  # copy list
            return

        need = k - len(curr)

        # pruning để tránh duyệt dư
        for j in range(start, n - (need - 1) + 1):
            curr.append(j)
            helper(j + 1, curr)
            curr.pop()

    helper(1, [])
    return res