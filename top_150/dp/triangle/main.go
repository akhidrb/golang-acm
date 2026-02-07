package main

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	k := ((n + 1) * n) / 2
	dp := make([]int, k)
	INF := int(^uint(0) >> 1)
	best := INF
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			index := (((i + 1) * i) / 2) + j
			rTopL := ((i * (i - 1)) / 2) + (j - 1)
			rTopR := ((i * (i - 1)) / 2) + j
			minParent := INF
			if j > 0 && j < len(triangle[i-1]) {
				minParent = min(dp[rTopL], dp[rTopR])
			} else if j > 0 {
				minParent = dp[rTopL]
			} else {
				minParent = dp[rTopR]
			}
			dp[index] = triangle[i][j] + minParent
			if i == n-1 && dp[index] < best {
				best = dp[index]
			}
		}
	}
	return best
}
