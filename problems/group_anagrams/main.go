package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagram := make(map[string][]string)
	for _, str := range strs {
		chars := []rune(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		if _, ok := anagram[string(chars)]; !ok {
			anagram[string(chars)] = []string{str}
		} else {
			anagram[string(chars)] = append(anagram[string(chars)], str)
		}
	}
	angrmList := make([][]string, 0)
	for _, val := range anagram {
		angrmList = append(angrmList, val)
	}
	return angrmList
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
