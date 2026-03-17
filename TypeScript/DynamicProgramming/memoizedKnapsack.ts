function memoizedKnapsack(
  W: number,
  wt: number[],
  val: number[],
  n: number
): number {

  // dp table initialized with -1
  const dp: number[][] = Array.from({ length: n }, () =>
    Array(W + 1).fill(-1)
  );

  function helper(index: number, remWeight: number): number {
    // base case
    if (index >= n || remWeight === 0) {
      return 0;
    }

    // return cached result
    if (dp[index][remWeight] !== -1) {
      return dp[index][remWeight];
    }

    // recursive case
    const exclude: number = helper(index + 1, remWeight);

    let include: number = 0;
    if (wt[index] <= remWeight) {
      include = val[index] + helper(index + 1, remWeight - wt[index]);
    }

    // store result
    dp[index][remWeight] = Math.max(exclude, include);
    return dp[index][remWeight];
  }

  return helper(0, W);
}