package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		p := []any{1, 2, 3, 4, 5}
		pTree := buildTree(p)
		output := diameterOfBinaryTree(pTree)
		assert.Equal(t, 3, output)
	})

	t.Run("2", func(t *testing.T) {
		p := []any{1, 2}
		pTree := buildTree(p)
		output := diameterOfBinaryTree(pTree)
		assert.Equal(t, 1, output)
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
