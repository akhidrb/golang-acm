package main

// [1,3], [6,9]
// [2,5]
// [1, 5], [6,9]

func insert(intervals [][]int, newInterval []int) [][]int {
	start, finish := -1, -1
	low, high := newInterval[0], newInterval[1]
	foundIntervals := make(map[int]bool)
	for i, _ := range intervals {
		if start == -1 && intervals[i][0] <= low && intervals[i][1] >= low {
			if intervals[i][0] < low {
				start = intervals[i][0]
			} else {
				start = low
			}
			if intervals[i][1] > high {
				finish = intervals[i][1]
			} else {
				finish = high
			}
			foundIntervals[i] = true
		}
		if intervals[i][1] >= high && intervals[i][1] <= high {
			foundIntervals[i] = true
			if intervals[i][0] < low {
				start = intervals[i][0]
			}
			if intervals[i][1] > high {
				finish = intervals[i][1]
			} else {
				finish = high
			}
		}
	}
	first := -1
	for i, _ := range foundIntervals {
		if first == -1 {
			first = i
		}
		intervals = append(intervals[:i], intervals[i+1:]...)
	}
	intervals = append(intervals[:first], append([][]int{{start, finish}}, intervals[first:]...)...)
	return intervals
}
