package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := inorderTraversal(constructTree())
		assert.Equal(t, []int{1, 3, 2}, res)
	})

	t.Run("2", func(t *testing.T) {
		res := inorderTraversal(nil)
		assert.Equal(t, []int{}, res)
	})

	t.Run("3", func(t *testing.T) {
		res := inorderTraversal(&TreeNode{Val: 1})
		assert.Equal(t, []int{1}, res)
	})

}

func constructTree() *TreeNode {
	val := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: nil,
		},
	}
	return val
}
