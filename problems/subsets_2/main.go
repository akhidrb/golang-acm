package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	cache := make(map[string]struct{})
	sort.Ints(nums)
	dfs(&res, []int{}, nums, 0, cache)
	return res
}

func dfs(res *[][]int, subs, nums []int, iter int, cache map[string]struct{}) {
	if iter >= len(nums) {
		dst := make([]int, len(subs))
		copy(dst, subs)
		if _, ok := cache[fmt.Sprint(dst)]; ok {
			return
		} else {
			*res = append(*res, dst)
			cache[fmt.Sprint(dst)] = struct{}{}
		}
		return
	}

	subs = append(subs, nums[iter])

	dfs(res, subs, nums, iter+1, cache)
	subs = subs[:len(subs)-1]
	dfs(res, subs, nums, iter+1, cache)
}
