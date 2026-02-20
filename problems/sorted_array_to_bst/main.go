package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return &TreeNode{}
	}
	node := &TreeNode{Val: nums[len(nums)/2]}
	if len(nums) >= 3 {
		node.Right = &TreeNode{}
		dfs(node.Right, nums[(len(nums)/2)+1:])
	}
	if len(nums) >= 2 {
		node.Left = &TreeNode{}
		dfs(node.Left, nums[:(len(nums)/2)])
	}
	return node
}

func dfs(node *TreeNode, sub []int) {
	if len(sub) == 0 {
		return
	}
	node.Val = sub[len(sub)/2]
	if len(sub) >= 3 {
		node.Right = &TreeNode{}
		dfs(node.Right, sub[(len(sub)/2)+1:])
	}
	if len(sub) >= 2 {
		node.Left = &TreeNode{}
		dfs(node.Left, sub[:(len(sub)/2)])
	}
}
