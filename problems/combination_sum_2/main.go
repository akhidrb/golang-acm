package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var dfs func(j int, sum int)
	res := make([][]int, 0)
	cur := make([]int, 0)
	dfs = func(j int, sum int) {
		if sum == target {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		for i := j; i < len(candidates); i++ {
			if i > j && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i]+sum <= target {
				cur = append(cur, candidates[i])
				dfs(i+1, sum+candidates[i])
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0, 0)
	return res
}
