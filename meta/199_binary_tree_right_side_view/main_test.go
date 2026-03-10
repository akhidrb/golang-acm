package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []any{1, 2, 3, nil, 5, nil, 4}
		root := buildTree(nums)
		res := rightSideView(root)
		exp := []int{1, 3, 4}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []any{1, 2, 3, 4, nil, nil, nil, 5}
		root := buildTree(nums)
		res := rightSideView(root)
		exp := []int{1, 3, 4, 5}
		assert.Equal(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []any{1, nil, 3}
		root := buildTree(nums)
		res := rightSideView(root)
		exp := []int{1, 3}
		assert.Equal(t, exp, res)
	})

	t.Run("4", func(t *testing.T) {
		nums := []any{}
		root := buildTree(nums)
		res := rightSideView(root)
		exp := []int{}
		assert.Equal(t, exp, res)
	})

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
