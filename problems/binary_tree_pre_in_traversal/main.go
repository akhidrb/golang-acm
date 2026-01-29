package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var myMap map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	myMap = make(map[int]int)
	for i, val := range inorder {
		myMap[val] = i
	}
	idx := 0
	return conTree(preorder, inorder, 0, len(inorder)-1, &idx)
}

func conTree(preorder []int, inorder []int, l, r int, idx *int) *TreeNode {
	if l > r {
		return nil
	}
	root := &TreeNode{Val: preorder[*idx]}
	*idx++
	mIdx := myMap[root.Val]
	root.Left = conTree(preorder, inorder, l, mIdx-1, idx)
	root.Right = conTree(preorder, inorder, mIdx+1, r, idx)
	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	binTree := buildTree(preorder, inorder)
	p := binTree
	printBin(p)
}

func printBin(p *TreeNode) {
	fmt.Println(p.Val)
	if p.Left != nil {
		printBin(p.Left)
	}
	if p.Right != nil {
		printBin(p.Right)
	}
}
