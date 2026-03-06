/*
N-Queens

The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

*/

function solveNQueens(n: number): string[][] {

    const res: string[][] = []
    const board: string[][] = Array.from({ length: n }, () => Array(n).fill('.'))

    function convertBoard(board: string[][]): string[] {
        return board.map(row => row.join(''))
    }

    function isValid(row: number, col: number): boolean {

        // check column
        for (let x = 0; x < row; x++) {
            if (board[x][col] === 'Q') return false
        }

        // top-left diagonal
        for (let r = row, c = col; r >= 0 && c >= 0; r--, c--) {
            if (board[r][c] === 'Q') return false
        }

        // top-right diagonal
        for (let r = row, c = col; r >= 0 && c < n; r--, c++) {
            if (board[r][c] === 'Q') return false
        }

        return true
    }

    function positionNextQueen(row: number) {

        if (row === n) {
            res.push(convertBoard(board))
            return
        }

        for (let col = 0; col < n; col++) {

            if (isValid(row, col)) {

                board[row][col] = 'Q'
                positionNextQueen(row + 1)
                board[row][col] = '.'

            }
        }
    }

    positionNextQueen(0)

    return res
}