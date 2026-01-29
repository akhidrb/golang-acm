package main

func lengthOfLongestSubstring(s string) int {
	lastIndex := make([]int, 128)
	i := 0
	longest := 0
	for j := 0; j < len(s); j++ {
		curr := s[j]
		if lastIndex[curr] > i {
			i = lastIndex[curr]
		}
		if j-i+1 > longest {
			longest = j - i + 1
		}
		lastIndex[curr] = j + 1
	}
	return longest
}
