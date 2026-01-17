package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	var dfs func(node *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftHeight := dfs(root.Left)
		rightHeight := dfs(root.Right)

		if leftHeight+rightHeight > diameter {
			diameter = leftHeight + rightHeight
		}
		return 1 + max(leftHeight, rightHeight)
	}
	dfs(root)
	return diameter
}
