package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		p := []any{3, 1, 4, nil, 2}
		k := 1
		pTree := buildTree(p)
		isSame := kthSmallest(pTree, k)
		assert.Equal(t, 1, isSame)
	})

	t.Run("2", func(t *testing.T) {
		p := []any{5, 3, 6, 2, 4, nil, nil, 1}
		k := 3
		pTree := buildTree(p)
		isSame := kthSmallest(pTree, k)
		assert.Equal(t, 3, isSame)
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
