function optimizedCombine(n: number, k: number): number[][] {
    const res: number[][] = [];

    function helper(start: number, curr: number[]) {
        if (curr.length === k) {
            res.push([...curr]); // clone array
            return;
        }

        const need = k - curr.length;

        for (let j = start; j <= n - (need - 1); j++) {
            curr.push(j);
            helper(j + 1, curr);
            curr.pop();
        }
    }

    helper(1, []);
    return res;
}