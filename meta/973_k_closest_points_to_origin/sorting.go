package main

func kClosestSorting(points [][]int, k int) [][]int {
	qs(points, k, 0, len(points)-1)
	return points[:k]
}

func qs(points [][]int, k, left, right int) {
	if left >= right {
		return
	}
	p := partition(points, left, right)
	if p == k {
		return
	} else if p < k-1 {
		qs(points, k, p+1, right)
	} else {
		qs(points, k, left, p-1)
	}
}

func partition(points [][]int, left, right int) int {
	i := left
	p := points[right]
	pivotSum := (p[0] * p[0]) + (p[1] * p[1])
	for j := left; j < right; j++ {
		x, y := points[j][0], points[j][1]
		sum := (x * x) + (y * y)
		if sum < pivotSum {
			points[j], points[i] = points[i], points[j]
			i++
		}
	}
	points[right], points[i] = points[i], points[right]
	return i
}
