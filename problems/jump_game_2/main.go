package main

func jump(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i]+i, nums[i-1])
	}
	ind := 0
	ans := 0
	for ind < len(nums)-1 {
		ans++
		ind = nums[ind]
	}
	return ans
}
