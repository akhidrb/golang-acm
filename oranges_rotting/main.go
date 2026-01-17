package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func orangesRotting(grid [][]int) int {
	dirs := [][]int{
		{0, 1},  // up
		{0, -1}, // down
		{1, 0},  // right
		{-1, 0}, // left
	}
	count := 0
	queue := make([][2]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	for len(queue) > 0 {
		for _, orange := range queue {
			queue = queue[1:]
			for i := range dirs {
				r, c := orange[0]+dirs[i][0], orange[1]+dirs[i][1]
				if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
					continue
				}
				if grid[r][c] != 1 {
					continue
				}
				grid[r][c] = 2
				queue = append(queue, [2]int{r, c})
			}
		}
		if len(queue) > 0 {
			count++
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return count
}
