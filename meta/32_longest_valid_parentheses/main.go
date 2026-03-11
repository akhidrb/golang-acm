package main

func longestValidParentheses(s string) int {
	stack := []int{-1}
	maxRes := 0
	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res := i - stack[len(stack)-1]
				if res > maxRes {
					maxRes = res
				}
			}
		}
	}

	return maxRes
}
