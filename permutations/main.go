package main

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	dfs(&res, nums, 0, n)
	return res
}

func dfs(res *[][]int, nums []int, i, n int) {
	if i == n {
		*res = append(*res, copyNums(nums))
		return
	}
	for j := i; j < n; j++ {
		nums[i], nums[j] = nums[j], nums[i]
		dfs(res, nums, i+1, n)
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func copyNums(nums []int) []int {
	c := make([]int, len(nums))
	copy(c, nums)
	return c
}
