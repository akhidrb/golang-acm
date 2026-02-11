package main

import (
	"strconv"
)

func isAdditiveNumber(num string) bool {
	res := false
	stack := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if res {
			return
		}
		if i == len(num) && len(stack) > 2 {
			res = true
			return
		}

		for j := i; j < len(num); j++ {
			if j > i && num[i] == '0' {
				break
			}
			part, _ := strconv.Atoi(num[i : j+1])
			n := len(stack)
			if n >= 2 {
				sum := stack[n-1] + stack[n-2]
				if part < sum {
					continue
				}
				if part > sum {
					break
				}
			}
			stack = append(stack, part)
			dfs(j + 1)
			stack = stack[:len(stack)-1]
		}
	}
	dfs(0)
	return res
}
