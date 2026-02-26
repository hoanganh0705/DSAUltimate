const uniqueSubset = (nums: number[]) => {
  nums.sort((a, b) => a - b);
  const res: number[][] = [];
  const helper = (i: number, curr: number[]) => {
    if (i === nums.length) {
      res.push([...curr]);
      return;
    }
    curr.push(nums[i]);
    helper(i + 1, curr);
    curr.pop();
    while (i < nums.length - 1 && nums[i] === nums[i + 1]) {
      i++;
    }
    helper(i + 1, curr);
  };
  helper(0, []);
  return res;
};