package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	type QueueItem struct {
		node *TreeNode
		col  int
	}
	queue := []QueueItem{{
		node: root,
		col:  0,
	}}
	minCol, maxCol := 0, 0
	columns := make(map[int][]int)
	columns[0] = []int{root.Val}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.node.Left != nil {
			queue = append(queue, QueueItem{
				node: node.node.Left,
				col:  node.col - 1,
			})
			columns[node.col-1] = append(columns[node.col-1], node.node.Left.Val)
			minCol = min(minCol, node.col-1)
		}
		if node.node.Right != nil {
			queue = append(queue, QueueItem{
				node: node.node.Right,
				col:  node.col + 1,
			})
			columns[node.col+1] = append(columns[node.col+1], node.node.Right.Val)
			maxCol = max(maxCol, node.col+1)
		}
	}

	res := make([][]int, maxCol-minCol+1)
	for i := minCol; i <= maxCol; i++ {
		res[i-minCol] = columns[i]
	}
	return res
}
