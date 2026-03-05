/*
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1–9 must occur exactly once in each row.

Each of the digits 1–9 must occur exactly once in each column.

Each of the digits 1–9 must occur exactly once in each of the 9 (3×3) sub-boxes of the grid.

The "." character indicates empty cells.

Example:
board = [

	["5","3",".",".","7",".",".",".","."],
	["6",".",".","1","9","5",".",".","."],
	[".","9","8",".",".",".",".","6","."],
	["8",".",".",".","6",".",".",".","3"],
	["4",".",".","8",".","3",".",".","1"],
	["7",".",".",".","2",".",".",".","6"],
	[".","6",".",".",".",".","2","8","."],
	[".",".",".","4","1","9",".",".","5"],
	[".",".",".",".","8",".",".","7","9"]

]

expected_solution = [

	["5","3","4","6","7","8","9","1","2"],
	["6","7","2","1","9","5","3","4","8"],
	["1","9","8","3","4","2","5","6","7"],
	["8","5","9","7","6","1","4","2","3"],
	["4","2","6","8","5","3","7","9","1"],
	["7","1","3","9","2","4","8","5","6"],
	["9","6","1","5","3","7","2","8","4"],
	["2","8","7","4","1","9","6","3","5"],
	["3","4","5","2","8","6","1","7","9"]

]
*/
package main

func solveSudoku(board [][]byte) {
	var isValid func(num byte, row int, col int) bool
	isValid = func(num byte, row int, col int) bool {
		for x := 0; x < 9; x++ {
			if board[row][x] == num {
				return false
			}
			if board[x][col] == num {
				return false
			}

			r := 3*(row/3) + x/3
			c := 3*(col/3) + x%3
			if board[r][c] == num {
				return false
			}
		}
		return true

	}

	var fillTheBoard func(board [][]byte) bool
	fillTheBoard = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					for num := byte('1'); num <= '9'; num++ {
						if isValid(num, i, j) {
							board[i][j] = num
							if fillTheBoard(board) {
								return true
							}
							board[i][j] = '.'
						}
					}
					return false
				}
			}
		}
		return true
	}
	fillTheBoard(board)
}
