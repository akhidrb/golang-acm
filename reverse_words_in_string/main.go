package main

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	strList := strings.Split(s, " ")
	i := 0
	for i < len(strList) {
		if strList[i] == "" {
			strList = append(strList[:i], strList[i+1:]...)
		} else {
			i++
		}

	}
	i = 0
	j := len(strList) - 1
	for i < j {
		strList[i], strList[j] = strings.Trim(strList[j], " "), strings.Trim(strList[i], " ")
		i++
		j--
	}
	return strings.Join(strList, " ")
}
