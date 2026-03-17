function canPartition(nums: number[]): boolean {
  const total = nums.reduce((a, b) => a + b, 0);

  if (total % 2 !== 0) return false;

  const target = total / 2;
  const n = nums.length;

  let prev: boolean[] = new Array(target + 1).fill(false);
  prev[0] = true;

  for (let i = 1; i <= n; i++) {
    const curr: boolean[] = new Array(target + 1).fill(false);
    curr[0] = true;

    for (let j = 1; j <= target; j++) {
      if (nums[i - 1] <= j) {
        curr[j] = prev[j - nums[i - 1]];
      }

      curr[j] = curr[j] || prev[j];
    }

    prev = curr;
  }

  return prev[target];
}