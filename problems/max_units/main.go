package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sum := 0
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for _, boxType := range boxTypes {
		if truckSize == 0 {
			break
		}
		numberOfBoxes := boxType[0]
		numberOfUnits := boxType[1]

		num := min(numberOfBoxes, truckSize)
		sum += num * numberOfUnits
		truckSize -= num
	}
	return sum
}
