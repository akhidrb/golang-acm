package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
		target, k := 5, 2
		numTree := buildTree(nums)
		numTarget := traverse(numTree, target)
		res := distanceK(numTree, numTarget, k)
		assert.Equal(t, 3, res)
	})

}

func traverse(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}
	left := traverse(node.Left, val)
	if left != nil {
		return left
	}
	return traverse(node.Right, val)
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
