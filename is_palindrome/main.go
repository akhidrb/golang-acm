package main

import (
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func clearString(str string) string {
	str = nonAlphanumericRegex.ReplaceAllString(str, "")
	return strings.ToLower(str)
}

func isPalindrome(s string) bool {
	s = clearString(s)
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
