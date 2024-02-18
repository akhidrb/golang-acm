package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt32-1, math.MaxInt32+1)
}

func dfs(p *TreeNode, minVal, maxVal int) bool {
	if p == nil {
		return true
	}
	if p.Val >= maxVal || p.Val <= minVal {
		return false
	}
	return dfs(p.Left, minVal, p.Val) && dfs(p.Right, p.Val, maxVal)
}
