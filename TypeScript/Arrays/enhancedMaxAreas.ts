function maxAreaOptimum(array: number[]): number {
    let left = 0;
    let right = array.length - 1;
    let maxArea = 0;

    while (left < right) {
        const area = Math.min(array[left], array[right]) * (right - left);
        maxArea = Math.max(maxArea, area);

        if (array[left] < array[right]) {
            left++;
        } else {
            right--;
        }
    }

    return maxArea;
}

console.log(maxAreaOptimum([3, 7, 5, 6, 8, 4])); // 21