package main

import "fmt"

func findPeakElement(nums []int) int {
	INF := -1 * ((1 << 31) + 1)
	nums = append([]int{INF}, nums...)
	nums = append(nums, INF)
	fmt.Println(INF)
	i := 1
	for nums[i] <= nums[i-1] || nums[i] <= nums[i+1] {
		i++
	}
	return i - 1
}
