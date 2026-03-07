package main

type OpenP struct {
	ch    rune
	index int
}

func minRemoveToMakeValid(s string) string {
	stack := make([]OpenP, 0, len(s))
	invalidIndex := make([]bool, len(s))
	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, OpenP{ch: ch, index: i})
		} else if ch == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				invalidIndex[i] = true
			}
		}
	}
	for _, val := range stack {
		index := val.index
		invalidIndex[index] = true
	}
	str := make([]rune, 0, len(s))
	for i, ch := range s {
		if !invalidIndex[i] {
			str = append(str, ch)
		}
	}

	return string(str)
}
