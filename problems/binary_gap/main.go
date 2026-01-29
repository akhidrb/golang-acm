package main

import (
	"fmt"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// Implement your solution here
	strA := fmt.Sprintf("%b", N)
	largest := 0
	first := false
	last := false
	num := 0
	for _, val := range strA {
		valInt := val - '0'
		if valInt == 1 {
			if !first {
				first = true
			} else {
				last = true
				first = true
			}
		}
		if last {
			if num > largest {
				largest = num
			}
			num = 0
			last = false
		}
		if valInt == 0 {
			num++
		}
	}
	return largest
}

func main() {
	A := 1041
	largest := Solution(A)
	fmt.Println(largest)
}
