package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		p := constructBinaryTree([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1})
		res := hasPathSum(p, 22)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		p := constructBinaryTree([]int{1, 2, 3})
		res := hasPathSum(p, 5)
		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		p := constructBinaryTree([]int{})
		res := hasPathSum(p, 0)
		assert.Equal(t, false, res)
	})
}

func constructBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
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
