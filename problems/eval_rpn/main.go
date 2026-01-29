package main

import "strconv"

type MyStack struct {
	stack []int
}

func Constructor() MyStack {
	stack := make([]int, 0)
	return MyStack{
		stack: stack,
	}
}

func (this *MyStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MyStack) Pop() int {
	n := len(this.stack)
	temp := this.stack[n-1]
	this.stack = this.stack[:n-1]
	return temp
}

func evalRPN(tokens []string) int {
	stack := Constructor()
	ops := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}
	for _, token := range tokens {
		if _, ok := ops[token]; ok {
			val1 := stack.Pop()
			val2 := stack.Pop()
			if token == "/" {
				val2 /= val1
			} else if token == "+" {
				val2 += val1
			} else if token == "-" {
				val2 -= val1
			} else if token == "*" {
				val2 *= val1
			}
			stack.Push(val2)
		} else {
			num, _ := strconv.Atoi(token)
			stack.Push(num)
		}
	}
	return stack.Pop()
}
