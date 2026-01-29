package main

import "math"

func maxProfit(prices []int) int {
	minVal, maxVal := math.MaxInt, 0

	for i := 0; i < len(prices); i++ {
		minVal = min(prices[i], minVal)
		maxVal = max(maxVal, prices[i]-minVal)
	}
	return max(maxVal, 0)
}
