def lengthOfLIS(nums):
    n = len(nums)

    def helper(curr_index, prev_index):
        # base case
        if curr_index >= n:
            return 0

        # exclude current element
        exclude = helper(curr_index + 1, prev_index)

        # include current element
        include = 0
        if prev_index == -1 or nums[curr_index] > nums[prev_index]:
            include = 1 + helper(curr_index + 1, curr_index)

        return max(exclude, include)

    return helper(0, -1)