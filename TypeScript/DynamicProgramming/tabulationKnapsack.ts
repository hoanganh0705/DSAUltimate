function knapSack(
  W: number,
  wt: number[],
  val: number[],
  n: number
): number {

  // dp[i][j] = max value using first i items with capacity j
  const dp: number[][] = Array.from({ length: n + 1 }, () =>
    Array(W + 1).fill(0)
  );

  for (let i = 1; i <= n; i++) {
    for (let j = 1; j <= W; j++) {

      const exclude: number = dp[i - 1][j];

      let include: number = 0;
      if (wt[i - 1] <= j) {
        include = val[i - 1] + dp[i - 1][j - wt[i - 1]];
      }

      dp[i][j] = Math.max(exclude, include);
    }
  }

  return dp[n][W];
}