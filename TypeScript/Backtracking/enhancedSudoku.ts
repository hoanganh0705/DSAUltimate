function enhancedSudoku(board: string[][]): void {
    const boxes: Array<Set<string>> = Array.from({ length: 9 }, () => new Set());
    const rows: Array<Set<string>> = Array.from({ length: 9 }, () => new Set());
    const cols: Array<Set<string>> = Array.from({ length: 9 }, () => new Set());

    const getBox = (row: number, col: number): number => {
        return Math.floor(row / 3) * 3 + Math.floor(col / 3);
    };

    for (let i = 0; i < 9; i++) {
        for (let j = 0; j < 9; j++) {
            if (board[i][j] !== ".") {
                const v = board[i][j];
                const b = getBox(i, j);
                boxes[b].add(v);
                rows[i].add(v);
                cols[j].add(v);
            }
        }
    }

    const backtrack = (r: number, c: number): boolean => {
        if (r === 9) return true;

        let nr = r;
        let nc = c + 1;

        if (nc === 9) {
            nr = r + 1;
            nc = 0;
        }

        if (board[r][c] !== ".") {
            return backtrack(nr, nc);
        }

        const b = getBox(r, c);

        for (let num = 1; num <= 9; num++) {
            const val = num.toString();

            if (boxes[b].has(val) || rows[r].has(val) || cols[c].has(val)) {
                continue;
            }

            board[r][c] = val;
            boxes[b].add(val);
            rows[r].add(val);
            cols[c].add(val);

            if (backtrack(nr, nc)) return true;

            boxes[b].delete(val);
            rows[r].delete(val);
            cols[c].delete(val);
            board[r][c] = ".";
        }

        return false;
    };

    backtrack(0, 0);
}