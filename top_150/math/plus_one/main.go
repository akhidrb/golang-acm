package main

func plusOne(digits []int) []int {
	index := len(digits) - 1
	for index >= 0 && digits[index] == 9 {
		digits[index] = 0
		index--
	}
	if index < 0 {
		digits = append(digits, 0)
		index = 0
	}
	digits[index] += 1
	return digits
}
