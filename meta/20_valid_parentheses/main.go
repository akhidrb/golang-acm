package main

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if ch == ')' && top != '(' {
				return false
			} else if ch == ']' && top != '[' {
				return false
			} else if ch == '}' && top != '{' {
				return false
			}
		} else {
			return false
		}

	}
	if len(stack) > 0 {
		return false
	}

	return true
}
