function findTargetSumWays(nums: number[], target: number): number {
  const n = nums.length;
  const summation = nums.reduce((a, b) => a + b, 0);

  const dp: (number | null)[][] = Array.from({ length: n }, () =>
    Array(2 * summation + 1).fill(null)
  );

  function helper(index: number, sumNums: number): number {
    if (index < 0) {
      return sumNums === target ? 1 : 0;
    }

    if (dp[index][sumNums + summation] !== null) {
      return dp[index][sumNums + summation] as number;
    }

    const negative = helper(index - 1, sumNums - nums[index]);
    const positive = helper(index - 1, sumNums + nums[index]);

    dp[index][sumNums + summation] = negative + positive;
    return dp[index][sumNums + summation] as number;
  }

  return helper(n - 1, 0);
}