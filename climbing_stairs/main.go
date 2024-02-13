package main

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	curr, prev := 1, 1
	for i := 2; i <= n; i++ {
		temp := curr
		curr += prev
		prev = temp
	}
	return curr
}
