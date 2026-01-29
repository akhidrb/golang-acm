package main

func climbStairs(n int, costs []int) int {
	memo := make([]int, n+1)
	const inf = int(^uint(0) >> 1)
	for i := range memo {
		memo[i] = inf
	}
	memo[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= 3; j++ {
			step := i - j
			if step < 0 || memo[step] == inf {
				continue
			}
			cost := memo[step] + costs[i-1] + (j * j)
			if cost < memo[i] {
				memo[i] = cost
			}
		}
	}
	return memo[n]
}
