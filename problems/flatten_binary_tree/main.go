package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	stack := make([]int, 0)
	dfs(root, &stack)
	p := root
	for len(stack) > 0 {
		p.Val = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		p.Left = nil
		if len(stack) > 0 {
			if p.Right == nil {
				p.Right = &TreeNode{}
			}
			p = p.Right
		}
	}
}

func dfs(node *TreeNode, stack *[]int) {
	if node == nil {
		return
	}
	dfs(node.Right, stack)
	dfs(node.Left, stack)
	*stack = append(*stack, node.Val)
}
