package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parent := make(map[*TreeNode]*TreeNode)
	buildParent(root, nil, parent)
	res := bfs(target, parent, k)
	return res
}

func bfs(target *TreeNode, parent map[*TreeNode]*TreeNode, k int) []int {
	queue := []*TreeNode{target}
	distance := 0
	visited := map[*TreeNode]bool{target: true}
	for len(queue) > 0 {
		size := len(queue)
		if distance == k {
			res := make([]int, 0)
			for _, node := range queue {
				res = append(res, node.Val)
			}
			return res
		}
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			neighbors := []*TreeNode{curr.Left, curr.Right, parent[curr]}
			for _, neighbor := range neighbors {
				if neighbor != nil && !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
		distance++
	}
	return []int{}
}

func buildParent(node, p *TreeNode, parent map[*TreeNode]*TreeNode) {
	if node == nil {
		return
	}
	parent[node] = p
	buildParent(node.Left, node, parent)
	buildParent(node.Right, node, parent)
}
