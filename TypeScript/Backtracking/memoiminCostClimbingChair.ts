const memoMinCostClimbingStairs = (cost: number[]) => {
    const n = cost.length;
    const minCost = new Array(n).fill(-1);

    const helper = (index: number) => {
        // base case
        if (index > n - 1) {
            return 0;
        }

        if (minCost[index] !== -1) {
            return minCost[index];
        }

        // one step
        const onestep = cost[index] + helper(index + 1);

        // two steps
        const twostep = cost[index] + helper(index + 2);

        minCost[index] = Math.min(onestep, twostep);

        return minCost[index];
    };

    return Math.min(helper(0), helper(1));
}