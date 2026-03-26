function maxProfit(prices: number[]): number {
    let left = 0;
    let profit = 0;

    for (let right = 1; right < prices.length; right++) {
        if (prices[right] < prices[left]) {
            left = right;
        } else {
            profit = Math.max(profit, prices[right] - prices[left]);
        }
    }

    return profit;
}