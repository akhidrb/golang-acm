package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	sum, mini, i, j := 0, 100000, 0, 0
	n := len(nums)
	for i < n && j < n {
		sum += nums[j]
		for sum >= target {
			mini = min(mini, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}
	return mini
}

func main() {
	fmt.Println(minSubArrayLen(6, []int{5, 2, 3, 5, 5}))
}