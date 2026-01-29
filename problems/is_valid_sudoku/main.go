package main

func isValidSudoku(board [][]byte) bool {
	rows := make(map[byte]bool)
	cols := make(map[byte]bool)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			val := board[i][j]
			if exists := rows[val]; exists && val != '.' {
				return false
			} else {
				rows[val] = true
			}

			val2 := board[j][i]
			if exists := cols[val2]; exists && val2 != '.' {
				return false
			} else {
				cols[val2] = true
			}
		}
		rows = make(map[byte]bool)
		cols = make(map[byte]bool)
	}

	for i := 0; i <= 6; i += 3 {
		for j := 0; j <= 6; j += 3 {
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					val := board[k][l]
					if exists := rows[val]; exists && val != '.' {
						return false
					} else {
						rows[val] = true
					}
				}
			}
			rows = make(map[byte]bool)
		}
	}

	return true
}
