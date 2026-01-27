package main

func generate(numRows int) [][]int {
	total := make([][]int, numRows)
	if numRows == 0 {
		return [][]int{}
	}
	total[0] = []int{1}
	if numRows == 1 {
		return total
	}
	total[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		nums := make([]int, i+1)
		nums[0] = 1
		nums[i] = 1
		for j := 1; j < i; j++ {
			nums[j] = total[i-1][j-1] + total[i-1][j]
		}
		total[i] = nums
	}
	return total
}
