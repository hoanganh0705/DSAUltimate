package main

func enhancedSudoku(board [][]byte) {
	boxes := make([]map[byte]bool, 9)
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		boxes[i] = map[byte]bool{}
		rows[i] = map[byte]bool{}
		cols[i] = map[byte]bool{}
	}

	getBox := func(row, col int) int {
		return (row/3)*3 + col/3
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				v := board[i][j]
				b := getBox(i, j)
				boxes[b][v] = true
				rows[i][v] = true
				cols[j][v] = true
			}
		}
	}

	var backtrack func(r, c int) bool
	backtrack = func(r, c int) bool {
		if r == 9 {
			return true
		}

		nr, nc := r, c+1
		if nc == 9 {
			nr = r + 1
			nc = 0
		}

		if board[r][c] != '.' {
			return backtrack(nr, nc)
		}

		b := getBox(r, c)

		for num := byte('1'); num <= '9'; num++ {
			if boxes[b][num] || rows[r][num] || cols[c][num] {
				continue
			}

			board[r][c] = num
			boxes[b][num] = true
			rows[r][num] = true
			cols[c][num] = true

			if backtrack(nr, nc) {
				return true
			}

			delete(boxes[b], num)
			delete(rows[r], num)
			delete(cols[c], num)
			board[r][c] = '.'
		}

		return false
	}

	backtrack(0, 0)
}
