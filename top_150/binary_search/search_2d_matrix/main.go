package main

func searchMatrix(matrix [][]int, target int) bool {
	start, end := 0, len(matrix)-1
	last := len(matrix[0]) - 1
	for start <= end {
		mid := (start + end) / 2
		if target == matrix[mid][last] {
			return true
		} else if target > matrix[mid][last] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	row := min(start, len(matrix)-1)
	start, end = 0, len(matrix[0])-1
	for start <= end {
		mid := (start + end) / 2
		if target == matrix[row][mid] {
			return true
		} else if target > matrix[row][mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
