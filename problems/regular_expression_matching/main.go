package main

func pop(stack []byte) byte {
	n := len(stack)
	top := stack[n-1]
	stack = stack[:n-1]
	return top
}

func dequeue(stack []byte) byte {
	first := stack[0]
	stack = stack[1:]
	return first
}

func push(stack []byte, val byte) {
	stack = append(stack, val)
}

func isMatch(s string, p string) bool {
	var prev byte
	i, j := 0, 0
	n, k := len(s), len(p)
	for i < n {
		if j < k {
			if p[j] != '*' {
				if j < k-1 && p[j+1] == '*' {

				} else if p[j] != s[i] && p[j] != '.' {
					return false
				}
				prev = p[j]
				if prev == s[i] || prev == '.' {
					i++
				}
			} else {
				for i < n && (prev == s[i] || prev == '.') {
					i++
				}
			}
			j++
		} else {
			return false
		}
	}

	if j != k {
		return false
	}

	return true
}
