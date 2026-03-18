function minDistance(word1: string, word2: string): number {
  const n: number = word1.length;
  const m: number = word2.length;

  function helper(i: number, j: number): number {
    // base cases
    if (i >= n && j >= m) return 0;
    if (i >= n) return m - j;
    if (j >= m) return n - i;

    if (word1[i] === word2[j]) {
      return helper(i + 1, j + 1);
    }

    const replace: number = 1 + helper(i + 1, j + 1);
    const del: number = 1 + helper(i + 1, j);
    const insert: number = 1 + helper(i, j + 1);

    return Math.min(replace, del, insert);
  }

  return helper(0, 0);
}