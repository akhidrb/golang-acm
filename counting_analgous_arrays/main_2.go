package main

func countAnalogousArrays2(consecutiveDifference []int32, lowerBound int32, upperBound int32) int32 {
	low, high := lowerBound, upperBound
	bounds := make([]int, upperBound-lowerBound+1)
	ind := 0
	for low < high {
		mid := (low + high) / 2
		sum := mid
		place := 1
		for _, val := range consecutiveDifference {
			sum -= val
			if sum > upperBound {
				place = 2
				break
			}
		}
		bounds[mid-lowerBound] = place
		if bounds[mid-lowerBound] == 2 && bounds[mid-lowerBound-1] == 1 {
			ind = int(mid)
			break
		} else if bounds[mid-lowerBound] == 1 && bounds[mid-lowerBound+1] == 2 {
			ind = int(mid + 1)
			break
		}
		if place == 2 {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if int32(ind) > lowerBound {
		return int32(ind) - lowerBound
	}
	return 0
}
