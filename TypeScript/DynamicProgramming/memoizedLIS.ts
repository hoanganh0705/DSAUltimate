function memoizedLIS(nums: number[]): number {
  const n: number = nums.length;

  const dp: number[][] = Array.from({ length: n }, () =>
    Array(n + 1).fill(-1)
  );

  function helper(curr: number, prev: number): number {
    if (curr >= n) return 0;

    if (dp[curr][prev + 1] !== -1) {
      return dp[curr][prev + 1];
    }

    const exclude: number = helper(curr + 1, prev);

    let include: number = 0;
    if (prev === -1 || nums[curr] > nums[prev]) {
      include = 1 + helper(curr + 1, curr);
    }

    dp[curr][prev + 1] = Math.max(exclude, include);
    return dp[curr][prev + 1];
  }

  return helper(0, -1);
}