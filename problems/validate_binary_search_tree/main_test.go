package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		p := constructBinaryTree([]int{2, 1, 3})
		res := isValidBST(p)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		p := constructBinaryTree([]int{5, 1, 4, 0, 0, 3, 6, 0, 0, 0, 0, 7, 8, 9})
		res := isValidBST(p)
		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		p := constructBinaryTree([]int{5, 4, 6, 0, 0, 3, 7})
		res := isValidBST(p)
		assert.Equal(t, false, res)
	})

	t.Run("4", func(t *testing.T) {
		p := constructBinaryTree([]int{2147483647, -2147483648})
		res := isValidBST(p)
		assert.Equal(t, true, res)
	})
}

func constructBinaryTree(nums []int) *TreeNode {
	root := &TreeNode{}
	p := root
	bfsTree(p, nums, 0)
	return root
}

func bfsTree(p *TreeNode, nums []int, index int) {
	p.Val = nums[index]
	if index*2+1 < len(nums) && nums[index*2+1] != 0 {
		p.Left = &TreeNode{}
		bfsTree(p.Left, nums, index*2+1)
	}
	if index*2+2 < len(nums) && nums[index*2+2] != 0 {
		p.Right = &TreeNode{}
		bfsTree(p.Right, nums, index*2+2)
	}
}
