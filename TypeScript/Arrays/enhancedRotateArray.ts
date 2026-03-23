function reverse(array: number[], start: number, end: number): number[] {
    while (start < end) {
        [array[start], array[end]] = [array[end], array[start]];
        start++;
        end--;
    }
    return array;
}

function rotateArray(array: number[], k: number): number[] {
    if (array.length === 0) {
        return [];
    }

    k = k % array.length;

    reverse(array, 0, array.length - 1);
    reverse(array, 0, k - 1);
    reverse(array, k, array.length - 1);

    return array;
}