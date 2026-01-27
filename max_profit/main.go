package main

func maxProfit(prices []int) int {
	totalProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			totalProfit += prices[i+1] - prices[i]
		}
	}
	return totalProfit
}
