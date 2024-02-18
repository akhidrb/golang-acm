package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, targetSum, 0)
}

func dfs(p *TreeNode, targetSum, sum int) bool {
	if p == nil {
		return false
	}
	if p.Left == nil && p.Right == nil && sum+p.Val == targetSum {
		return true
	}
	return dfs(p.Left, targetSum, sum+p.Val) || dfs(p.Right, targetSum, sum+p.Val)
}
