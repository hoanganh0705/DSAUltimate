def spaceOptimizedTabulationKnapsack(W, wt, val, n):
    prev = [0] * (W + 1)
    curr = [0] * (W + 1)

    for i in range(1, n + 1):
        for j in range(1, W + 1):
            exclude = prev[j]

            include = 0
            if wt[i - 1] <= j:
                include = val[i - 1] + prev[j - wt[i - 1]]

            curr[j] = max(exclude, include)

        # move current row to previous
        prev = curr[:]

    return prev[W]
