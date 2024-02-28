package main

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	dfs(&res, candidates, []int{}, 0, target, 0)
	return res
}

func dfs(res *[][]int, candidates, nums []int, ind, target, sum int) {
	if sum == target {
		*res = append(*res, copyNums(nums))
		return
	}
	if ind >= len(candidates) || sum > target {
		return
	}
	sum += candidates[ind]
	nums = append(nums, candidates[ind])
	dfs(res, candidates, nums, ind, target, sum)
	sum -= candidates[ind]
	nums = nums[:len(nums)-1]
	dfs(res, candidates, nums, ind+1, target, sum)
}

func copyNums(nums []int) []int {
	c := make([]int, len(nums))
	copy(c, nums)
	return c
}
