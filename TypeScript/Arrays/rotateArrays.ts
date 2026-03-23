function rotateArray(array: number[], k: number): number[] {
    if (array.length === 0) {
        return [];
    }

    k = k % array.length;

    const temp = array.slice(-k);

    for (let i = array.length - k - 1; i >= 0; i--) {
        array[i + k] = array[i];
    }

    for (let i = 0; i < temp.length; i++) {
        array[i] = temp[i];
    }

    return array;
}
