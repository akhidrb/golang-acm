package main

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	dfs(&res, nums, 0)
	return res
}

func dfs(res *[][]int, nums []int, i int) {
	if i == len(nums) {
		*res = append(*res, copyNums(nums))
		return
	}
	for j := i; j < len(nums); j++ {
		nums[i], nums[j] = nums[j], nums[i]
		dfs(res, nums, i+1)
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func copyNums(nums []int) []int {
	c := make([]int, len(nums))
	copy(c, nums)
	return c
}
