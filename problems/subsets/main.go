package main

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	dfs(&res, []int{}, nums, 0)
	return res
}

func dfs(res *[][]int, subs, nums []int, iter int) {
	if iter >= len(nums) {
		dst := make([]int, len(subs))
		copy(dst, subs)
		*res = append(*res, dst)
		return
	}

	subs = append(subs, nums[iter])

	dfs(res, subs, nums, iter+1)
	subs = subs[:len(subs)-1]
	dfs(res, subs, nums, iter+1)
}
