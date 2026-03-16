package main

func largestIsland(grid [][]int) int {
	var dfs func(i, j int)
	area := 0
	id := 2
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
			return
		}
		area++
		grid[i][j] = id
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	maxArea := 0
	areas := make(map[int]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area = 0
				dfs(i, j)
				areas[id] = area
				id++
				maxArea = max(maxArea, area)
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				seen := make(map[int]bool)
				area = 1
				dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
				for _, dir := range dirs {
					x, y := i+dir[0], j+dir[1]
					if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
						continue
					}
					id = grid[x][y]
					if !seen[id] {
						seen[id] = true
						area += areas[id]
					}
				}

				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}
