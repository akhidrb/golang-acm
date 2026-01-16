package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	output := make([]int, 0)
	pot(root, &output)
	return output
}

func pot(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	pot(root.Left, output)
	pot(root.Right, output)
	*output = append(*output, root.Val)
}
