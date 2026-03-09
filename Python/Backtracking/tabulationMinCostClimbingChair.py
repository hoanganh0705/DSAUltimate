def tabulationMinCostClimbingStairs(cost):
    # write code here
    n = len(cost)

    mincost = [0] * (n + 1)

    mincost[0] = 0
    mincost[1] = 0

    for i in range(2, n + 1):
        onestep = cost[i - 1] + mincost[i - 1]
        twostep = cost[i - 2] + mincost[i - 2]

        mincost[i] = min(onestep, twostep)

    return mincost[n]