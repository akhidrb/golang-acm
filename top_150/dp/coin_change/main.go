package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	const INF = int(^uint(0) >> 1)
	for a := 1; a <= amount; a++ {
		dp[a] = INF
		for _, val := range coins {
			num := a - val
			if num >= 0 && dp[num] != INF {
				dp[a] = min(dp[num]+1, dp[a])
			}
		}

	}
	if dp[amount] == INF {
		return -1
	}
	return dp[amount]
}
