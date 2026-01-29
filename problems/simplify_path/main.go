package main

import (
	"fmt"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func simplifyPath(path string) string {
	if len(path) == 0 {
		return "/"
	}
	pathList := strings.Split(path, "/")
	myStack := make(Stack, 0)
	for i, _ := range pathList {
		if pathList[i] != "" && pathList[i] != "." {
			myStack.Push(pathList[i])
		}
	}
	strBuild := ""
	pop := 0
	for len(myStack) > 0 {
		str, _ := myStack.Pop()
		if str != ".." {
			if pop > 0 {
				pop--
			} else {
				strBuild = "/" + str + strBuild
			}
		} else {
			pop++
		}
	}
	if strBuild == "" {
		return "/"
	}
	return strBuild
}

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
