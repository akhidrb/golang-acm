package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	dfs(root, &sum, false)
	return sum
}

func dfs(root *TreeNode, sum *int, isLeft bool) {
	if root == nil {
		return
	}
	dfs(root.Left, sum, true)
	if root.Left == nil && root.Right == nil && isLeft {
		*sum += root.Val
	}
	dfs(root.Right, sum, false)
}
