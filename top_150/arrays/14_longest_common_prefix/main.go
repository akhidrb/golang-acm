package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	prefixes := make(map[string]int)
	for i := 0; i < len(strs[0]); i++ {
		prefix := strs[0][:i+1]
		prefixes[prefix] = 1
	}

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(strs[i]); j++ {
			prefix := strs[i][:j+1]
			if _, ok := prefixes[prefix]; ok {
				prefixes[prefix]++
			}
		}
	}
	highest := 0
	highestStr := ""
	for key, value := range prefixes {
		if value == len(strs) {
			if len(key) > highest {
				highest = len(key)
				highestStr = key
			}
		}
	}

	return highestStr
}
