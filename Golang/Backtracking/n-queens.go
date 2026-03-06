/*
N-Queens

The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.


Example:
board =
	[
	['.','Q','.','.'],
	['.','.','.','Q'],
	['Q','.','.','.'],
	['.','.','Q','.']
	]

res =
[
 [
  ".Q..",
  "...Q",
  "Q...",
  "..Q."
 ],

 [
  "..Q.",
  "Q...",
  "...Q",
  ".Q.."
 ]
]
*/

package main

func solveNQueens(n int) [][]string {
	res := [][]string{}
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	var convertBoard func(board [][]byte) []string
	convertBoard = func(board [][]byte) []string {
		res := make([]string, len(board))

		for i := range board {
			res[i] = string(board[i])
		}

		return res
	}

	var isValid func(row int, col int) bool
	isValid = func(row int, col int) bool {
		for x := 0; x < row; x++ {
			if board[x][col] == 'Q' {
				return false
			}
		}

		// top-left diagonal
		for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
			if board[r][c] == 'Q' {
				return false
			}
		}

		// top-right diagonal
		for r, c := row-1, col+1; r >= 0 && c < n; r, c = r-1, c+1 {
			if board[r][c] == 'Q' {
				return false
			}
		}
		return true
	}

	var positionNextQueen func(row int)
	positionNextQueen = func(row int) {
		if row == n {
			res = append(res, convertBoard(board))
			return
		}

		for col := 0; col < n; col++ {
			if isValid(row, col) {
				board[row][col] = 'Q'
				positionNextQueen(row + 1)
				board[row][col] = '.'
			}
		}
	}

	positionNextQueen(0)
	return res
}
