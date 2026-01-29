package main

func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	ans := 0
	i, j := 1, n-2
	maxLeft, maxRight := height[0], height[n-1]
	for i <= j {
		if maxLeft < maxRight {
			if height[i] >= maxLeft {
				maxLeft = height[i]
			} else {
				ans += maxLeft - height[i]
			}
			i++
		} else {
			if height[j] >= maxRight {
				maxRight = height[j]
			} else {
				ans += maxRight - height[j]
			}
			j--
		}
	}
	return ans
}
