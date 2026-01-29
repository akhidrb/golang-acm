package main

import "fmt"

func maxArea(height []int) int {
	maxRes := 0
	i, j := 0, len(height)-1
	for i < j {
		var result int
		if height[i] < height[j] {
			result = height[i] * (j - i)
			i++
		} else {
			result = height[j] * (j - i)
			j--
		}
		if result > maxRes {
			maxRes = result
		}
	}
	return maxRes
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
