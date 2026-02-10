package main

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(counter map[int]int)
	curr := make([]int, 0)
	counter := make(map[int]int)
	for _, val := range nums {
		counter[val]++
	}
	dfs = func(counter map[int]int) {
		if len(curr) == len(nums) {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		for k, v := range counter {
			if v > 0 {
				counter[k]--
				curr = append(curr, k)
				dfs(counter)
				counter[k]++
				curr = curr[:len(curr)-1]
			}

		}

	}
	dfs(counter)
	return res
}
