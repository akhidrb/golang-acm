package main

func minWindow(s string, t string) string {
	countMap := make([]int, 128)
	for _, ch := range t {
		countMap[ch]++
	}
	count, start, end := len(t), 0, 0
	startInd, minLen := 0, 1000000

	for end < len(s) {
		if countMap[s[end]] > 0 {
			count--
		}
		countMap[s[end]]--
		end++
		for count == 0 {
			if end-start < minLen {
				startInd = start
				minLen = end - start
			}
			if countMap[s[start]] == 0 {
				count++
			}
			countMap[s[start]]++
			start++
		}
	}

	if minLen == 1000000 {
		return ""
	}
	return s[startInd : startInd+minLen]
}
