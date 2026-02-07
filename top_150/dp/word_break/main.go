package main

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]struct{})
	maxLen := 0
	for _, val := range wordDict {
		wordSet[val] = struct{}{}
		if len(val) > maxLen {
			maxLen = len(val)
		}
	}
	dict := make([]bool, len(s)+1)
	j := 0
	dict[j] = true
	for i := 1; i <= len(s); i++ {
		start := max(i-maxLen, 0)
		for j := start; j < i; j++ {
			m := s[j:i]
			if _, ok := wordSet[m]; ok && dict[j] {
				dict[i] = true
				break
			}
		}
	}
	return dict[len(s)]
}
