def threeSum(nums):
    nums.sort()
    res = []

    for i in range(len(nums)):
        if i == 0 or nums[i] != nums[i - 1]:
            left = i + 1
            right = len(nums) - 1

            while left < right:
                sumThree = nums[i] + nums[left] + nums[right]

                if sumThree == 0:
                    res.append([nums[i], nums[left], nums[right]])
                    left += 1

                    while left < right and nums[left] == nums[left - 1]:
                        left += 1

                elif sumThree < 0:
                    left += 1
                else:
                    right -= 1

    return res


def threeSumNew(nums):
    res = set()

    for i in range(len(nums)):
        need = set()

        for j in range(i + 1, len(nums)):
            valueNeeded = -(nums[i] + nums[j])

            if valueNeeded in need:
                triplet = tuple(sorted((nums[i], nums[j], valueNeeded)))
                res.add(triplet)

            need.add(nums[j])

    return [list(triplet) for triplet in res]