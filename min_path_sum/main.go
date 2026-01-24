package main

func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dp := make([]int, (row*col)+1)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			index := (i * col) + j
			val := grid[i][j]
			up, left := 100000, 100000
			if i > 0 {
				up = dp[((i-1)*col)+j]
			}
			if j > 0 {
				left = dp[(i*col)+(j-1)]
			}
			if i > 0 || j > 0 {
				dp[index] = min(up, left) + val
			} else {
				dp[index] = val
			}
		}
	}

	return dp[row*col-1]
}
