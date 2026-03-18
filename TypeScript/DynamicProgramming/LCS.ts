function longestCommonSubsequence(
  text1: string,
  text2: string
): number {

  const n: number = text1.length;
  const m: number = text2.length;

  function helper(index1: number, index2: number): number {
    // base case
    if (index1 >= n || index2 >= m) {
      return 0;
    }

    // match
    if (text1[index1] === text2[index2]) {
      return 1 + helper(index1 + 1, index2 + 1);
    }

    // no match
    return Math.max(
      helper(index1, index2 + 1),
      helper(index1 + 1, index2)
    );
  }

  return helper(0, 0);
}