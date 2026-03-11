package main

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	minVal := math.MaxInt
	maxVal := 0
	for i := 1; i < len(prices); i++ {
		minVal = min(minVal, prices[i])
		maxVal = max(maxVal, prices[i]-minVal)
	}
	return max(maxVal, 0)
}
