package main

import (
	"strconv"
	"strings"
)

const (
	digit   = "digit"
	op      = "op"
	bracket = "bracket"
)

func calculate(s string) int {
	total := 0
	st := make(Stack, 0)
	i := len(s) - 1
	var sb strings.Builder
	for i < len(s) {
		ch := string(s[i])
		charType := getCharType(ch)
		if charType == op {
			st.Push(ch)
		} else if charType == digit {
			for i < len(s) && charType == digit {
				sb.WriteString(ch)
				i++
				if i < len(s) {
					ch = string(s[i])
					charType = getCharType(ch)
				}
			}
			st.Push(sb.String())
			sb.Reset()
			continue
		} else if charType == bracket {
			if ch == "(" {
				st.Push(ch)
			} else {
				total += handleStack(st)
			}
		}
		i++
	}
	total += handleStack(st)
	return total
}

func handleStack(st Stack) int {
	total := 0
	for !st.IsEmpty() {
		val, _ := st.Pop()
		if val == "(" {
			return total
		}
		if val == "-" {
			numV, exists := st.Pop()
			var num int
			if exists {
				num, _ = strconv.Atoi(numV)
			}
			total = num - total
		}
		if num, err := strconv.Atoi(val); err == nil {
			total += num
		}
	}
	return total
}

func getCharType(ch string) string {
	_, err := strconv.Atoi(ch)
	if err == nil {
		return digit
	}
	switch ch {
	case "(", ")":
		return bracket
	case "+", "-":
		return op
	}
	return ""
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Top() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element, true
	}
}
