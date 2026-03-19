function tabulationLIS(nums: number[]): number {
  const n: number = nums.length;

  const dp: number[][] = Array.from({ length: n + 1 }, () =>
    Array(n + 1).fill(0)
  );

  for (let i = n - 1; i >= 0; i--) {
    for (let j = i; j >= 0; j--) {

      const exclude: number = dp[i + 1][j];

      let include: number = 0;
      if (j === 0 || nums[i] > nums[j - 1]) {
        include = 1 + dp[i + 1][i + 1];
      }

      dp[i][j] = Math.max(exclude, include);
    }
  }

  return dp[0][0];
}