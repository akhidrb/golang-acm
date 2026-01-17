package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	maxHeight := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 1
		}
		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)
		if leftHeight > maxHeight {
			maxHeight = leftHeight
		}
		if rightHeight > maxHeight {
			maxHeight = rightHeight
		}
		return 1 + max(leftHeight, rightHeight)
	}
	dfs(root)
	return maxHeight
}
