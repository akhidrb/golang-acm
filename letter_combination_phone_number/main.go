package main

import (
	"strconv"
)

var (
	m = map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ans := make([]string, 0)
	bt(digits, 0, []byte{}, &ans)
	return ans
}

func bt(digits string, index int, str []byte, ans *[]string) {
	if index == len(digits) {
		*ans = append(*ans, string(str))
		return
	}

	num, _ := strconv.Atoi(string(digits[index]))
	for i, _ := range m[num] {
		str = append(str, m[num][i])
		bt(digits, index+1, str, ans)
		str = str[:len(str)-1]
	}

}
