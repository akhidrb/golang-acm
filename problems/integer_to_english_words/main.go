package main

import (
	"strings"
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	below20 := []string{
		"", "One", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve",
		"Thirteen", "Fourteen", "Fifteen", "Sixteen",
		"Seventeen", "Eighteen", "Nineteen",
	}

	tens := []string{
		"", "", "Twenty", "Thirty", "Forty",
		"Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
	}

	thousands := []string{"", "Thousand", "Million", "Billion"}

	var helper func(int) string
	helper = func(n int) string {
		if n == 0 {
			return ""
		} else if n < 20 {
			return below20[n] + " "
		} else if n < 100 {
			return tens[n/10] + " " + helper(n%10)
		} else {
			return below20[n/100] + " Hundred " + helper(n%100)
		}
	}

	result := ""
	i := 0

	for num > 0 {
		if num%1000 != 0 {
			result = helper(num%1000) + thousands[i] + " " + result
		}
		num /= 1000
		i++
	}

	return strings.TrimSpace(result)
}
