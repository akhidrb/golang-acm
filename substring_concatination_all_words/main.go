package main

func findSubstring(s string, words []string) []int {
	m := make(map[string]int)
	for _, val := range words {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	wLen := len(words[0])
	window := wLen * len(words)

	i := 0
	inds := make([]int, 0)
	for i+window <= len(s) {
		str := s[i : window+i]
		p := make(map[string]int)
		for key, value := range m {
			p[key] = value
		}
		if checkWindow(str, p, wLen) {
			inds = append(inds, i)
		}
		i++
	}

	return inds
}

func checkWindow(str string, p map[string]int, wLen int) bool {
	i := 0
	for i+wLen <= len(str) {
		sub := str[i : wLen+i]
		if _, ok := p[sub]; ok {
			if p[sub] == 0 {
				return false
			}
			p[sub]--
		} else {
			return false
		}
		i += wLen
	}
	return true
}
