package main

func maxSubArray(nums []int) int {
	highest := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(sum+nums[i], nums[i])
		if sum > highest {
			highest = sum
		}
	}
	return highest
}
