package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	str := make([]byte, 0)
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if (s[i] <= 'z' && s[i] >= 'a') || (s[i] >= '0' && s[i] <= '9') {
			str = append(str, s[i])
		}
	}
	s = string(str)
	fmt.Println(s)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
