package main

func numIslands(grid [][]byte) int {
	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				count++
				dfs(grid, r, c)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, r, c int) {
	dirs := [][]int{
		{1, 0},  // down
		{-1, 0}, // up
		{0, 1},  // right
		{0, -1}, // left
	}
	k, j := len(grid)-1, len(grid[0])-1
	if r < 0 || c < 0 || r > k || c > j || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	for _, d := range dirs {
		dfs(grid, r+d[0], c+d[1])
	}
}
