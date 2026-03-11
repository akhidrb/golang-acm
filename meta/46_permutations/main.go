package main

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
		}

		for j := i; j < len(nums); j++ {
			nums[j], nums[i] = nums[i], nums[j]
			dfs(i + 1)
			nums[j], nums[i] = nums[i], nums[j]
		}

	}
	dfs(0)
	return res
}
