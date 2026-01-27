package main

func productExceptSelf(nums []int) []int {
	answers := make([]int, len(nums))
	answers[0] = 1
	for i := 1; i < len(nums); i++ {
		answers[i] = answers[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		answers[i] *= nums[i+1]
		nums[i] *= nums[i+1]
	}
	return answers
}
