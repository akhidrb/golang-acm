package main

import "fmt"

func removeDuplicates(nums []int) int {
	index, repeat := 1, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			repeat++
		} else {
			repeat = 0
		}
		if repeat < 2 {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
