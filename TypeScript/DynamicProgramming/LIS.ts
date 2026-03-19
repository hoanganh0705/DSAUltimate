function lengthOfLIS(nums: number[]): number {
  const n: number = nums.length;

  function helper(currIndex: number, prevIndex: number): number {
    // base case
    if (currIndex >= n) {
      return 0;
    }

    // exclude
    const exclude: number = helper(currIndex + 1, prevIndex);

    // include
    let include: number = 0;
    if (
      prevIndex === -1 ||
      nums[currIndex] > nums[prevIndex]
    ) {
      include = 1 + helper(currIndex + 1, currIndex);
    }

    return Math.max(exclude, include);
  }

  return helper(0, -1);
}