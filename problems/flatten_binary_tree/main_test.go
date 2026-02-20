package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		p := []any{1, 2, 5, 3, 4, nil, 6}
		pTree := buildTree(p)
		flatten(pTree)
		exp := []any{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6}
		expTree := buildTree(exp)
		assert.Equal(t, expTree, pTree)
	})

	t.Run("2", func(t *testing.T) {
		p := []any{}
		pTree := buildTree(p)
		flatten(pTree)
		exp := []any{}
		expTree := buildTree(exp)
		assert.Equal(t, expTree, pTree)
	})

	t.Run("3", func(t *testing.T) {
		p := []any{0}
		pTree := buildTree(p)
		flatten(pTree)
		exp := []any{0}
		expTree := buildTree(exp)
		assert.Equal(t, expTree, pTree)
	})

	t.Run("4", func(t *testing.T) {
		p := []any{1, nil, 2, 3}
		pTree := buildTree(p)
		flatten(pTree)
		exp := []any{1, nil, 2, nil, 3}
		expTree := buildTree(exp)
		assert.Equal(t, expTree, pTree)
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
