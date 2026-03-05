'''
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
'''

def solveSudoku(board):
    # The function modifies the board in place to present the solution. Hence there's no need to return the board
    def isValid(num, row, col):
        for x in range(9):
            # Check row
            if board[row][x] == num:
                return False
            # Check col
            if board[x][col] == num:
                return False
            # box check
            r = 3 * (row // 3) + x//3
            c = 3* (col//3) + x%3
            if board[r][c] == num:
                return False
        return True
 
    def fillTheBoard(board):
        # identify any empty cell
        for row in range (9):
            for col in range (9):
                if board[row][col] == ".":
                    for num in '123456789':
                        if isValid(num, row, col):
                            board[row][col] = num
                            if (fillTheBoard(board)):
                                return True
                            board[row][col] = '.' # backtracking step
                    return False
        return True

    fillTheBoard(board)