package main

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	res := make([][]int, 0, n+1)
	i := 0
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	res = append(res, newInterval)
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}
