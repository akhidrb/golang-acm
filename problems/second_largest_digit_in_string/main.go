package main

import (
	"regexp"
	"strconv"
)

func secondHighest(s string) int {
	re := regexp.MustCompile("[^0-9]+")
	nums := re.ReplaceAllString(s, "")
	largest := -1
	second := -1
	for _, val := range nums {
		num, _ := strconv.Atoi(string(val))
		if num > largest {
			second = largest
			largest = num
		}
		if num > second && num != largest {
			second = num
		}
	}
	if second == largest {
		return -1
	}
	return second
}
