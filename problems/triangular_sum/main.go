package main

func triangularSum(nums []int) int {
	last := len(nums) - 1
	for last > 0 {
		for i := 0; i < last; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		last--
	}
	return nums[0]
}
