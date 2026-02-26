function permuteUnique(nums: number[]): number[][] {
    const result: number[][] = [];
    
    function backtrack(start: number) {
        if (start === nums.length - 1) {
            result.push([...nums]);
            return;
        }
        
        const used = new Set<number>();
        for (let i = start; i < nums.length; i++) {
            if (used.has(nums[i])) {
                continue;
            }
            used.add(nums[i]);
            [nums[start], nums[i]] = [nums[i], nums[start]]; // Swap
            backtrack(start + 1);
            [nums[start], nums[i]] = [nums[i], nums[start]]; // Backtrack (swap back)
        }
    }
    
    backtrack(0);
    return result;
}

// Example usage:
console.log(permuteUnique([1, 4])); // Output: [[1,4],[4,1]]
console.log(permuteUnique([1, 4, 5])); // Output: [[1,4,5],[1,5,4],[4,1,5],[4,5,1],[5,1,4],[5,4,1]]