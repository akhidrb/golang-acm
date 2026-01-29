package stack

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	str := strings.Replace(s, " ", "", -1)
	total, _ := calcRecur(str, 0)
	return total
}

func calcRecur(s string, i int) (int, int) {
	var res int
	sign := 1
	var currNum int
	for i < len(s) {
		if isDig(string(s[i])) {
			currNum, i = getDigits(s, i)
			res += sign * currNum
			continue
		}
		if s[i] == '+' {
			sign = 1
			i++
			continue
		}
		if s[i] == '-' {
			sign = -1
			i++
			continue
		}
		if s[i] == '(' {
			currNum, i = calcRecur(s, i+1)
			res += sign * currNum
			continue
		}
		if s[i] == ')' {
			i++
			break
		}
	}
	return res, i
}

func getDigits(s string, i int) (int, int) {
	res := 0
	for i < len(s) && isDig(string(s[i])) {
		val, _ := strconv.Atoi(string(s[i]))
		res = (res * 10) + val
		i++
	}
	return res, i
}

func isDig(val string) bool {
	if _, err := strconv.Atoi(val); err != nil {
		return false
	}
	return true
}
