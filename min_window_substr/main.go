package main

import "fmt"

func minWindow(s string, t string) string {
	i, j := 0, 0
	isAll, all := 0, len(t)
	tMap := make(map[string]int)
	smallest := len(s) + 1
	for _, val := range t {
		if _, ok := tMap[string(val)]; !ok {
			tMap[string(val)] = 1
		} else {
			tMap[string(val)]++
		}
	}
	subStr := ""
	for j < len(s) {
		if val, ok := tMap[string(s[j])]; ok {
			tMap[string(s[j])]--
			if val > 0 {
				isAll++
			}
		}
		for isAll == all {
			size := j - i + 1
			if size < smallest {
				smallest = size
				subStr = s[i : j+1]
			}
			if val, ok := tMap[string(s[i])]; ok {
				tMap[string(s[i])]++
				if val == 0 {
					isAll--
				}
			}
			i++
		}
		j++
	}
	return subStr
}

func main() {
	fmt.Println(minWindow("bba", "ab"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
}
