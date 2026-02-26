const subset = (nums: number[]) => {
  const output: number[][] = [];
  const helper = (i: number, subset: number[]) => {
    if (i === nums.length) {
      output.push([...subset]);
      return;
    }
    helper(i + 1, subset);
    subset.push(nums[i]);
    helper(i + 1, subset);
    subset.pop();
  };
  helper(0, []);
  return output;
};