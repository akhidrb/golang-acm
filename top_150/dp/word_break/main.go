package main

func dp(s string, wordSet map[string]struct{}) bool {
	dict := make([]bool, len(s)+1)
	j := 0
	dict[j] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			m := s[j:i]
			if _, ok := wordSet[m]; ok && dict[j] {
				dict[i] = true
			}
		}
	}
	return dict[len(s)]
}

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]struct{})
	for _, val := range wordDict {
		wordSet[val] = struct{}{}
	}
	return dp(s, wordSet)
}
