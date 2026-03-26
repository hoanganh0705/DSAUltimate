function threeSum(nums: number[]): number[][] {
    nums.sort((a, b) => a - b);
    const res: number[][] = [];

    for (let i = 0; i < nums.length; i++) {
        if (i === 0 || nums[i] !== nums[i - 1]) {
            let left = i + 1;
            let right = nums.length - 1;

            while (left < right) {
                const sumThree = nums[i] + nums[left] + nums[right];

                if (sumThree === 0) {
                    res.push([nums[i], nums[left], nums[right]]);
                    left++;

                    while (left < right && nums[left] === nums[left - 1]) {
                        left++;
                    }

                } else if (sumThree < 0) {
                    left++;
                } else {
                    right--;
                }
            }
        }
    }

    return res;
}


function threeSumNew(nums: number[]): number[][] {
    const res = new Set<string>();

    for (let i = 0; i < nums.length; i++) {
        const need = new Set<number>();

        for (let j = i + 1; j < nums.length; j++) {
            const valueNeeded = -(nums[i] + nums[j]);

            if (need.has(valueNeeded)) {
                const triplet = [nums[i], nums[j], valueNeeded].sort((a, b) => a - b);
                res.add(triplet.join(","));
            }

            need.add(nums[j]);
        }
    }

    return Array.from(res).map(str => str.split(",").map(Number));
}