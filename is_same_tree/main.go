package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	stack := []*TreeNode{p}
	stackQ := []*TreeNode{q}

	for len(stack) > 0 && len(stackQ) > 0 {
		node := stack[len(stack)-1]
		nodeQ := stackQ[len(stackQ)-1]
		stack = stack[:len(stack)-1]
		stackQ = stackQ[:len(stackQ)-1]
		if node.Val != nodeQ.Val {
			return false
		}
		sum := 0
		if node.Right != nil {
			stack = append(stack, node.Right)
			sum++
		}
		if nodeQ.Right != nil {
			stackQ = append(stackQ, nodeQ.Right)
			sum++
		}
		if sum == 1 {
			return false
		}
		sum = 0
		if node.Left != nil {
			stack = append(stack, node.Left)
			sum++
		}
		if nodeQ.Left != nil {
			stackQ = append(stackQ, nodeQ.Left)
			sum++
		}
		if sum == 1 {
			return false
		}
	}
	if len(stack) > 0 || len(stackQ) > 0 {
		return false
	}
	return true
}

func buildTree(level []any) *TreeNode {
	if len(level) == 0 || level[0] == nil {
		return nil
	}

	root := &TreeNode{Val: level[0].(int)}
	q := []*TreeNode{root}
	i := 1

	for len(q) > 0 && i < len(level) {
		node := q[0]
		q = q[1:]

		// left child
		if i < len(level) && level[i] != nil {
			node.Left = &TreeNode{Val: level[i].(int)}
			q = append(q, node.Left)
		}
		i++

		// right child
		if i < len(level) && level[i] != nil {
			node.Right = &TreeNode{Val: level[i].(int)}
			q = append(q, node.Right)
		}
		i++
	}

	return root
}
