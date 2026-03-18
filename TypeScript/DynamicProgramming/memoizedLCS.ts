function memoizedLCS(text1: string, text2: string): number {
  const n: number = text1.length;
  const m: number = text2.length;

  const dp: number[][] = Array.from({ length: n }, () =>
    Array(m).fill(-1)
  );

  function helper(i: number, j: number): number {
    // base case
    if (i >= n || j >= m) {
      return 0;
    }

    if (dp[i][j] !== -1) {
      return dp[i][j];
    }

    if (text1[i] === text2[j]) {
      dp[i][j] = 1 + helper(i + 1, j + 1);
    } else {
      dp[i][j] = Math.max(
        helper(i, j + 1),
        helper(i + 1, j)
      );
    }

    return dp[i][j];
  }

  return helper(0, 0);
}