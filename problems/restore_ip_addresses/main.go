package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	parts := make([]string, 0, 4)
	var dfs func(i int)
	dfs = func(i int) {
		if len(parts) == 4 {
			if i >= len(s) {
				res = append(res, strings.Join(parts, "."))
			}
			return
		}
		for j := 1; j < 4; j++ {
			if i+j > len(s) {
				continue
			}
			part := s[i : i+j]
			if len(part) > 1 && part[0] == '0' {
				continue
			}
			num, _ := strconv.Atoi(part)
			if num < 256 {
				parts = append(parts, part)
				dfs(i + j)
				parts = parts[:len(parts)-1]
			}
		}
	}
	dfs(0)
	return res
}
