package main

func removeDuplicateLetters(s string) string {
	seen := make(map[rune]int)
	lastOcc := make(map[rune]int)
	stack := make([]rune, 0, 26)

	for i, ch := range s {
		lastOcc[ch] = i
	}

	for i, ch := range s {
		if seen[ch-'a'] == 0 {
			for len(stack) > 0 && ch < stack[len(stack)-1] && i < lastOcc[stack[len(stack)-1]] {
				seen[stack[len(stack)-1]-'a'] = 0
				stack = stack[:len(stack)-1]
			}
			seen[ch-'a'] = 1
			stack = append(stack, ch)
		}
	}
	return string(stack)
}
