package main

func partition(s string) [][]string {
	res := make([][]string, 0)
	subs := make([]string, 0)
	pal := make([][]bool, len(s))
	for i := range pal {
		pal[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 2 || pal[i+1][j-1]) {
				pal[i][j] = true
			}
		}
	}
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(s) {
			tmp := append([]string(nil), subs...)
			res = append(res, tmp)
			return
		}
		for j := 1; j+i <= len(s); j++ {
			part := s[i : j+i]
			if pal[i][j+i-1] {
				subs = append(subs, part)
				dfs(i + j)
				subs = subs[:len(subs)-1]
			}
		}
	}
	dfs(0)
	return res
}
