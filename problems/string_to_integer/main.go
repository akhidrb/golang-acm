package main

import (
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	op := 1
	if s[0] == '-' {
		op = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	inf := (1 << 31) - 1
	num := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			break
		}
		num = (num * 10) + int(ch)
		if num > inf {
			break
		}
	}
	num = op * num
	if num >= 0 {
		if num > inf {
			num = inf
		}
	} else {
		inf = -1 * (1 << 31)
		if num < inf {
			num = inf
		}
	}
	return num
}
