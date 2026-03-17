function spaceOptimizedTabulationKnapsack(
  W: number,
  wt: number[],
  val: number[],
  n: number
): number {

  let prev: number[] = new Array(W + 1).fill(0);
  let curr: number[] = new Array(W + 1).fill(0);

  for (let i = 1; i <= n; i++) {
    for (let j = 1; j <= W; j++) {

      const exclude: number = prev[j];

      let include: number = 0;
      if (wt[i - 1] <= j) {
        include = val[i - 1] + prev[j - wt[i - 1]];
      }

      curr[j] = Math.max(exclude, include);
    }

    // copy current to previous
    prev = [...curr];
  }

  return prev[W];
}
