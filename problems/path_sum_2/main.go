package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	total [][]int
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	total = make([][]int, 0)
	if root == nil {
		return total
	}
	dfs(root, targetSum, root.Val, []int{root.Val})
	return total
}

func dfs(p *TreeNode, targetSum, sum int, nums []int) {
	if p.Left == nil && p.Right == nil && sum == targetSum {
		nl := make([]int, len(nums))
		copy(nl, nums)
		total = append(total, nl)
	}
	if p.Left != nil {
		dfs(p.Left, targetSum, sum+p.Left.Val, append(nums, p.Left.Val))
	}
	if p.Right != nil {
		dfs(p.Right, targetSum, sum+p.Right.Val, append(nums, p.Right.Val))
	}
}
