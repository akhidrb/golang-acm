package main

import "strconv"

func sequentialDigits(low int, high int) []int {
	s := strconv.Itoa(low)
	first, _ := strconv.Atoi(string(s[0]))
	nums := make([]int, 0)
	n := len(s)
	if n-1+first > 9 {
		first = 1
		n++
	}
	for i := 1; i < n; i++ {
		mod := (first % 10) + 1
		first *= 10
		first += mod
	}
	str := strconv.Itoa(first)
	lastNum, _ := strconv.Atoi(string(str[len(str)-1]))
	for first <= high && len(str) > 1 {
		if first >= low {
			nums = append(nums, first)
		}
		if lastNum == 9 {
			first = 1
			for i := 1; i < len(str)+1; i++ {
				mod := (first % 10) + 1
				first *= 10
				first += mod
			}
			str = strconv.Itoa(first)
			if first <= high {
				nums = append(nums, first)
			}
		}
		if lastNum == 9 && len(str) == 9 {
			break
		}
		str = str[1:]
		last := str[len(str)-1]
		first, _ = strconv.Atoi(str)
		lastNum, _ = strconv.Atoi(string(last))
		first *= 10
		lastNum++
		first += lastNum
		str = strconv.Itoa(first)
	}

	return nums
}
