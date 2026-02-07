package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, m*n)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := (i * n) + j
			top := ((i - 1) * n) + j
			left := (i * n) + j - 1
			if obstacleGrid[i][j] == 0 {
				if i > 0 && j > 0 {
					dp[k] = dp[top] + dp[left]
				} else if j > 0 {
					dp[k] = dp[left]
				} else if i > 0 {
					dp[k] = dp[top]
				}
			}
		}
	}
	return dp[m*n-1]
}
