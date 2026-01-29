package main

import "fmt"

func canJump(nums []int) bool {
	reachable := 0
	for i, val := range nums {
		if i > reachable {
			return false
		}
		reachable = max(reachable, val+i)
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{3, 0, 8, 2, 0, 0, 1}))
}
