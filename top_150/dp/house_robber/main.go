package main

func rob(nums []int) int {
	dp := make([]int, len(nums))
	best := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
		for j := 0; j < i-1; j++ {
			dp[i] = max(dp[i], dp[j]+nums[i])
		}
		if dp[i] > best {
			best = dp[i]
		}
	}

	return best
}
