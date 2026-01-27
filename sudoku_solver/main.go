package main

func solveSudoku(board [][]byte) {
	backtrack(board)
}

func backtrack(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for num := byte('1'); num <= byte('9'); num++ {
					if isValid(board, i, j, num) {
						board[i][j] = num
						if backtrack(board) {
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

func isValid(board [][]byte, r, c int, num byte) bool {
	r0 := 3 * (r / 3)
	c0 := 3 * (c / 3)
	for i := 0; i < 9; i++ {
		if board[r][i] == num {
			return false
		}
		if board[i][c] == num {
			return false
		}
		if board[r0+i/3][c0+i%3] == num {
			return false
		}
	}
	return true
}
