package main

import (
	"strings"
)

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, word := range paths {
		if word == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if len(word) > 0 && word != "." {
			stack = append(stack, word)
		}
	}
	path = strings.Join(stack, "/")
	path = "/" + path
	return path
}
