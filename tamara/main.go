package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack []uint32

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val uint32) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (uint32, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Top() (uint32, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element, true
	}
}

func Solution(S string) int {
	opList := strings.Split(S, " ")
	numStack := make(Stack, 0)
	for _, op := range opList {
		if num := action(op, &numStack); num == -1 {
			return -1
		}
	}
	val, ok := numStack.Pop()
	if !ok {
		return -1
	}
	return int(val)
}

func action(op string, numStack *Stack) int {
	const largest = (1 << 20) - 1
	switch op {
	case "POP":
		if _, ok := numStack.Pop(); !ok {
			return -1
		}
		break
	case "DUP":
		if val, ok := numStack.Top(); ok {
			numStack.Push(val)
		} else {
			return -1
		}
		break
	case "+":
		val1, ok1 := numStack.Pop()
		val2, ok2 := numStack.Pop()
		if !ok1 || !ok2 {
			return -1
		}
		if val1+val2 > largest {
			return -1
		} else {
			numStack.Push(val1 + val2)
		}
		break
	case "-":
		val1, ok1 := numStack.Pop()
		val2, ok2 := numStack.Pop()
		if !ok1 || !ok2 {
			return -1
		}
		if val1 >= val2 {
			numStack.Push(val1 - val2)
		} else {
			return -1
		}
		break
	default:
		i, err := strconv.Atoi(op)
		if err != nil || i > largest || i < 0 {
			return -1
		}
		numStack.Push(uint32(i))
		return 0
	}
	return 0
}

func main() {
	fmt.Println(Solution("4 5 6 - 7 +"))
	fmt.Println(Solution("13 DUP 4 POP 5 DUP + DUP + -"))
	fmt.Println(Solution("5 6 + -"))
	fmt.Println(Solution("3 DUP 5 - -"))
	fmt.Println(Solution("1048575 DUP +"))
}
