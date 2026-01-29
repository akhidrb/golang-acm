package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	stack := []*TreeNode{root}
	stack2 := []*TreeNode{root}
	for len(stack) > 0 && len(stack2) > 0 {
		node := stack[len(stack)-1]
		node2 := stack2[len(stack2)-1]
		stack = stack[:len(stack)-1]
		stack2 = stack2[:len(stack2)-1]
		if node.Val != node2.Val {
			return false
		}

		sum := 0
		if node.Right != nil {
			stack = append(stack, node.Right)
			sum++
		}
		if node2.Left != nil {
			stack2 = append(stack2, node2.Left)
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
		if node2.Right != nil {
			stack2 = append(stack2, node2.Right)
			sum++
		}
		if sum == 1 {
			return false
		}
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
