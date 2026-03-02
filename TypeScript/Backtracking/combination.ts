function combine(n: number, k: number): number[][] {
    const res: number[][] = [];

    function helper(start: number, curr: number[]) {
        if (curr.length === k) {
            res.push([...curr]); // clone array
            return;
        }

        for (let j = start; j <= n; j++) {
            curr.push(j);
            helper(j + 1, curr);
            curr.pop();
        }
    }

    helper(1, []);
    return res;
}