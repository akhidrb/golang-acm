package main

import "fmt"

func same[T comparable](x, y T) bool {
	if x == y {
		return true
	}
	return false
}

func main() {
	fmt.Println(same("hello", "yes"))
	fmt.Println(same("hello", "hello"))
	fmt.Println(same(1, 2))
	fmt.Println(same(1.2, 1.2))
}
