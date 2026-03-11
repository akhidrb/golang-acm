package main

func maxArea(height []int) int {
	maxRes := 0
	i, j := 0, len(height)-1
	for i < j {
		var res int
		if height[i] < height[j] {
			res = height[i] * (j - i)
			i++
		} else {
			res = height[j] * (j - i)
			j--
		}
		if res > maxRes {
			maxRes = res
		}
	}
	return maxRes
}
