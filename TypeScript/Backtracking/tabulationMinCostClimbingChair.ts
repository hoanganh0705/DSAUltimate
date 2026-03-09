function tabulationMinCostClimbingStairs(cost: number[]): number {

    const n = cost.length
    const mincost: number[] = new Array(n + 1).fill(0)

    mincost[0] = 0
    mincost[1] = 0

    for (let i = 2; i <= n; i++) {

        const oneStep = cost[i - 1] + mincost[i - 1]
        const twoStep = cost[i - 2] + mincost[i - 2]

        mincost[i] = Math.min(oneStep, twoStep)
    }

    return mincost[n]
}