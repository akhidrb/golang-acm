package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	tMap := make(map[string]int)
	for _, val := range magazine {
		if _, ok := tMap[string(val)]; !ok {
			tMap[string(val)] = 1
		} else {
			tMap[string(val)]++
		}
	}
	for _, val := range ransomNote {
		if mapVal, ok := tMap[string(val)]; ok {
			if mapVal > 0 {
				tMap[string(val)]--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
