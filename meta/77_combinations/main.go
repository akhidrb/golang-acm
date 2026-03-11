package main

import "fmt"

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	nums := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if len(nums) == k {
			tmp := make([]int, k)
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}
		for j := i; j <= n; j++ {
			nums = append(nums, j)
			dfs(j + 1)
			nums = nums[:len(nums)-1]
		}
	}
	dfs(1)
	fmt.Println(res)
	return res
}
