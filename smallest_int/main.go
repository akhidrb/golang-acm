package main

import "fmt"

func Solution(A []int) int {
	largest := 0
	for _, val := range A {
		if val > largest {
			largest = val
		}
	}
	newA := make([]int, largest+1)
	for _, val := range A {
		if val > 0 {
			newA[val] = 1
		}
	}
	for i := 1; i < largest; i++ {
		if newA[i] == 0 {
			return i
		}
	}
	return largest + 1
}

func main() {
	fmt.Println(Solution([]int{-1, -3}))
}
