package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nums := make([]int, 0)
	dfs(root, &nums)
	return nums
}

func dfs(p *TreeNode, nums *[]int) {
	if p == nil {
		return
	}
	dfs(p.Left, nums)
	*nums = append(*nums, p.Val)
	dfs(p.Right, nums)
}
