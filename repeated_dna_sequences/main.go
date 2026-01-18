package main

func findRepeatedDnaSequences(s string) []string {
	countMap := make(map[string]int)
	res := make([]string, 0)
	start, end := 0, 9
	for end < len(s) {
		countMap[s[start:end+1]]++
		if countMap[s[start:end+1]] == 2 {
			res = append(res, s[start:end+1])
		}
		start++
		end++
	}
	return res
}
