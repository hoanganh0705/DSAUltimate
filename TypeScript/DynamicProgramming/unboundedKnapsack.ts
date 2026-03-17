function unboundedKnapsack(
  N: number,
  W: number,
  val: number[],
  wt: number[]
): number {

  const dp: number[][] = Array.from({ length: N + 1 }, () =>
    Array(W + 1).fill(0)
  );

  for (let i = 1; i <= N; i++) {
    for (let j = 1; j <= W; j++) {

      const exclude: number = dp[i - 1][j];

      let include: number = 0;
      if (wt[i - 1] <= j) {
        include = val[i - 1] + dp[i][j - wt[i - 1]];
      }

      dp[i][j] = Math.max(include, exclude);
    }
  }

  return dp[N][W];
}