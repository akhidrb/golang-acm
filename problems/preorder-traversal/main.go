package main

import "fmt"

type TreeNode struct {
	Value    int
	Children []TreeNode
}

// Function to traverse the tree in pre-order
func preOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Value, " ")
	for _, v := range root.Children {
		preOrderTraversal(&v)
	}
}

func main() {
	// Create a tree
	root := &TreeNode{
		Value: 10,
		Children: []TreeNode{
			{
				Value: 20,
				Children: []TreeNode{
					{
						Value: 30,
						Children: []TreeNode{
							{
								Value: 40,
								Children: []TreeNode{
									{
										Value: 50,
									},
								},
							},
						},
					},
					{
						Value: 60,
					},
				},
			},
			{
				Value: 70,
				Children: []TreeNode{
					{
						Value: 80,
					},
				},
			},
		},
	}

	// Call the pre-order traversal function
	fmt.Print("Pre-order traversal: ")
	preOrderTraversal(root)
	fmt.Println()
}
