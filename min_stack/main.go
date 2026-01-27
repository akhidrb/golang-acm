package main

type MinStack struct {
	stack []int
	minS  []int
}

func Constructor() MinStack {
	stack := make([]int, 0)
	minS := make([]int, 0)
	return MinStack{
		stack: stack,
		minS:  minS,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	n := len(this.minS)
	if n > 0 {
		top := this.minS[n-1]
		this.minS = append(this.minS, min(val, top))
	} else {
		this.minS = append(this.minS, val)
	}
}

func (this *MinStack) Pop() {
	n := len(this.stack)
	this.stack = this.stack[:n-1]
	this.minS = this.minS[:n-1]
}

func (this *MinStack) Top() int {
	n := len(this.stack)
	return this.stack[n-1]
}

func (this *MinStack) GetMin() int {
	n := len(this.minS)
	return this.minS[n-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
