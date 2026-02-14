package main

func numberOfWays(s string) int64 {
	var right0, right1 int64 = 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			right0++
		} else {
			right1++
		}
	}
	var sum int64 = 0
	var left0, left1 int64 = 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			right0--
			sum += left1 * right1
			left0++
		} else {
			right1--
			sum += left0 * right0
			left1++
		}
	}
	return sum
}
