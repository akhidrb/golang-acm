package main

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	combs := make([]int, 0, k)
	var dfs func(num, sum int)
	dfs = func(num, sum int) {
		if len(combs) == k {
			if sum == n {
				tmp := append([]int(nil), combs...)
				res = append(res, tmp)
			}
			return
		}

		for j := num; j <= 9; j++ {
			if j+sum > n {
				break
			}
			combs = append(combs, j)
			dfs(j+1, sum+j)
			combs = combs[:len(combs)-1]
		}
	}
	dfs(1, 0)
	return res
}
