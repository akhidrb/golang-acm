package main

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	st := make(Stack, 0)
	dSt := make(Stack, 0)

	i := 0
	for i < len(s) {
		for i < len(s) && s[i] != ')' {
			if _, err := strconv.Atoi(string(s[i])); err == nil {
				var sb strings.Builder
				isDig := true
				for i < len(s) && isDig {
					sb.WriteString(string(s[i]))
					i++
					if i < len(s) {
						if _, err = strconv.Atoi(string(s[i])); err != nil {
							isDig = false
						}
					}
				}
				st.Push(sb.String())
			} else {
				st.Push(string(s[i]))
				i++
			}

		}
		for !st.IsEmpty() && st.Top() != "(" {
			top, _ := st.Pop()
			if _, err := strconv.Atoi(top); err == nil || top == "+" || top == "-" {
				dSt.Push(top)
			}
		}
		if st.Top() == "(" {
			st.Pop()
		}
		num := 0
		for !dSt.IsEmpty() {
			top, _ := dSt.Pop()
			if top == "-" {
				val, _ := dSt.Pop()
				if dig, err := strconv.Atoi(val); err == nil {
					num -= dig
				}
			}
			if dig, err := strconv.Atoi(top); err == nil {
				num += dig
			}
		}
		st.Push(strconv.Itoa(num))
		i++
	}
	for !st.IsEmpty() {
		top, _ := st.Pop()
		if _, err := strconv.Atoi(top); err == nil || top == "+" || top == "-" {
			dSt.Push(top)
		}
	}
	total := 0
	for !dSt.IsEmpty() {
		top, _ := dSt.Pop()
		if top == "-" {
			val, _ := dSt.Pop()
			if dig, err := strconv.Atoi(val); err == nil {
				total -= dig
			}
		}
		if dig, err := strconv.Atoi(top); err == nil {
			total += dig
		}
	}

	return total
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

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element
	}
}
