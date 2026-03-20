function twoCitySchedCost(costs: number[][]): number {
    costs.sort((a, b) => (a[0] - a[1]) - (b[0] - b[1]));

    const n = costs.length;
    let cost = 0;

    for (let i = 0; i < n / 2; i++) {
        cost += costs[i][0];
    }

    for (let i = n / 2; i < n; i++) {
        cost += costs[i][1];
    }

    return cost;
}