package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	i := 0
	longest := 0
	iter := 0
	for i < len(s) {
		if val, ok := m[rune(s[i])]; ok {
			i = val + 1
			m = make(map[rune]int)
			iter = 0
		}
		m[rune(s[i])] = i
		iter++
		i++
		if iter > longest {
			longest = iter
		}
	}
	return longest
}
