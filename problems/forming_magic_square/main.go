package main

import "math"

var (
	magic = [][][]int32{
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
	}
)

func formingMagicSquare(s [][]int32) int32 {
	minDiff := 2 << 9
	for _, row := range magic {
		diffs := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				diff := s[i][j] - row[i][j]
				diffs += int(math.Abs(float64(diff)))
			}
		}
		if diffs < minDiff {
			minDiff = diffs
		}
	}

	return int32(minDiff)
}
