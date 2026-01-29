package main

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	i := 0
	for i < len(s) {
		val := romanMap[s[i]]
		if i+1 < len(s) {
			valAfter := romanMap[s[i+1]]
			if valAfter > val {
				sum += valAfter - val
				i += 2
			} else {
				sum += val
				i++
			}
		} else {
			sum += val
			i++
		}
	}

	return sum
}
