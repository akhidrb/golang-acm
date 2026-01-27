package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == target {
		return root
	}
	// Return non-nil return value from the recursive calls
	left := dfs(root.Left, target)
	if left != nil {
		return left
	}
	// At this point, we know left is nil, and right could be nil or non-nil
	// We return the right child's recursive call result directly because
	// - If it's non-nil, we should return it
	// - If it's nil, then both left and right are nil, and we want to return nil
	return dfs(root.Right, target)
}
