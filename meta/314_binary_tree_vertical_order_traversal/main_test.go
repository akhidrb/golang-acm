package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []any{3, 9, 20, nil, nil, 15, 7}
		root := buildTree(nums)
		res := verticalOrder(root)
		exp := [][]int{{9}, {3, 15}, {20}, {7}}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []any{3, 9, 8, 4, 0, 1, 7}
		root := buildTree(nums)
		res := verticalOrder(root)
		exp := [][]int{{4}, {9}, {3, 0, 1}, {8}, {7}}
		assert.Equal(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []any{1, 2, 3, 4, 10, 9, 11, nil, 5, nil, nil, nil, nil, nil, nil, nil, 6}
		root := buildTree(nums)
		res := verticalOrder(root)
		exp := [][]int{{4}, {2, 5}, {1, 10, 9, 6}, {3}, {11}}
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
