package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	return bfs(root)
}

func bfs(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	queue := []*TreeNode{node}
	res := make([]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			if i == size-1 {
				res = append(res, top.Val)
			}
			queue = queue[1:]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
	}
	return res
}
