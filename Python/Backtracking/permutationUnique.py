def permute(nums):
    def backtrack(start=0):
        # If we've reached the end of the array, we found a permutation
        if start == len(nums):
            result.append(nums[:])
            return
        
        hash = {}
        for i in range(start, len(nums)):

            # Skip duplicates
            if nums[i] in hash:
                continue
            hash[nums[i]] = True

            # Swap the current element with the start element
            nums[start], nums[i] = nums[i], nums[start]
            # Recurse with the next element as the start
            backtrack(start + 1)
            # Backtrack: swap back to the original configuration
            nums[start], nums[i] = nums[i], nums[start]
    
    result = []
    backtrack()
    return result