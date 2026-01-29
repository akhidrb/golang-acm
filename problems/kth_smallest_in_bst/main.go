package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	output := make([]int, 0)
	val := inorderTraverse(root, &output, k)
	return val
}

func inorderTraverse(root *TreeNode, output *[]int, k int) int {
	if root == nil {
		return -1
	}
	if val := inorderTraverse(root.Left, output, k); val != -1 {
		return val
	}
	*output = append(*output, root.Val)
	if len(*output) == k {
		return root.Val
	}
	if val := inorderTraverse(root.Right, output, k); val != -1 {
		return val
	}
	return -1
}
