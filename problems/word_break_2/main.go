package main

import "strings"

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]bool)
	for _, val := range wordDict {
		wordMap[val] = true
	}
	cur := make([]string, 0)
	res := make([]string, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(s) {
			res = append(res, strings.Join(cur, " "))
			return
		}
		for j := i; j < len(s); j++ {
			m := s[i : j+1]
			if wordMap[m] {
				cur = append(cur, m)
				dfs(j + 1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0)
	return res
}
