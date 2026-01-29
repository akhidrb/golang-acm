package main

func majorityElement(nums []int) int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	highestKey, highestCount := 0, 0
	for key, value := range counter {
		if value > highestCount {
			highestKey = key
			highestCount = value
		}
	}
	return highestKey
}
