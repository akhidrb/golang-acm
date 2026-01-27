package main

func totalNQueens(n int) int {
	sum := 0
	for r := 0; r < n; r++ {
		grid := make([][]bool, n)
		for i := 0; i < n; i++ {
			grid[i] = make([]bool, n)
		}
		num := dfs(grid, 0, r, n)
		if num == n {
			sum++
		}
	}
	return sum
}

func dfs(grid [][]bool, i, j, n int) int {
	if !(i >= 0 && i < n && j >= 0 && j < n) {
		return 0
	}
	if grid[i][j] {
		return 0
	}
	r, c := i, j
	for j := 0; j < n; j++ {
		grid[i][j] = true
	}
	for i := 0; i < n; i++ {
		grid[i][j] = true
	}
	i, j = r, c
	for i > 0 && j > 0 {
		i--
		j--
		grid[i][j] = true
	}
	i, j = r, c
	for i > 0 && j < n-1 {
		i--
		j++
		grid[i][j] = true
	}
	i, j = r, c
	for i < n-2 && j < n-1 {
		i++
		j++
		grid[i][j] = true
	}
	i, j = r, c
	for i < n-2 && j > 0 {
		i++
		j--
		grid[i][j] = true
	}
	i, j = r, c
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !grid[i][j] {
				r, c = i, j
				return 1 + dfs(grid, r, c, n)
			}
		}
	}
	return 1
}
