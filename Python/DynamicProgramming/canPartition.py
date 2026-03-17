def canPartition(nums):
    total = sum(nums)

    if total % 2 != 0:
        return False

    target = total // 2
    n = len(nums)

    prev = [False] * (target + 1)
    prev[0] = True

    for i in range(1, n + 1):
        curr = [False] * (target + 1)
        curr[0] = True

        for j in range(1, target + 1):
            # pick
            if nums[i - 1] <= j:
                curr[j] = prev[j - nums[i - 1]]

            # don't pick
            curr[j] = curr[j] or prev[j]

        prev = curr

    return prev[target]