package main

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	dfs(&res, []int{}, 1, k, n)
	return res
}

func dfs(res *[][]int, nums []int, ind, k, n int) {
	if k == 0 {
		c := make([]int, len(nums))
		copy(c, nums)
		*res = append(*res, c)
		return
	}
	if ind > n {
		return
	}
	nums = append(nums, ind)
	dfs(res, nums, ind+1, k-1, n)
	nums = nums[:len(nums)-1]
	dfs(res, nums, ind+1, k, n)
}
