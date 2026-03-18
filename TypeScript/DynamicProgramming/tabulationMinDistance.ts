function tabulationMinDistance(word1: string, word2: string): number {
  const n: number = word1.length;
  const m: number = word2.length;

  const dp: number[][] = Array.from({ length: n + 1 }, () =>
    Array(m + 1).fill(0)
  );

  // base cases
  for (let j = 0; j <= m; j++) {
    dp[0][j] = j;
  }

  for (let i = 0; i <= n; i++) {
    dp[i][0] = i;
  }

  // fill dp
  for (let i = 1; i <= n; i++) {
    for (let j = 1; j <= m; j++) {

      if (word1[i - 1] === word2[j - 1]) {
        dp[i][j] = dp[i - 1][j - 1];
      } else {
        const replace: number = 1 + dp[i - 1][j - 1];
        const del: number = 1 + dp[i - 1][j];
        const insert: number = 1 + dp[i][j - 1];

        dp[i][j] = Math.min(replace, del, insert);
      }

    }
  }

  return dp[n][m];
}