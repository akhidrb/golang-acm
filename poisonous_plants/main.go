package main

func poisonousPlants(p []int32) int32 {
	arrPrev, arr := p, make([]int32, 0)
	days := int32(0)
	if len(p) == 1 {
		return days
	}
	for len(arrPrev) > len(arr) {
		if len(arr) > 0 {
			arrPrev = arr
		}
		arr = []int32{arrPrev[0]}
		for i := 1; i < len(arrPrev); i++ {
			if arrPrev[i] <= arrPrev[i-1] {
				arr = append(arr, arrPrev[i])
			}
		}
		days++
	}

	return days - 1
}
