package main

import "strconv"

func validWordAbbreviation(word string, abbr string) bool {
	i, j := 0, 0
	for i < len(abbr) && j < len(word) {
		ch := abbr[i]
		if ch <= 'z' && ch >= 'a' {
			if ch != word[j] {
				return false
			}
			i++
			j++
			continue
		}
		if ch == '0' {
			return false
		}
		num := ""
		for i < len(abbr) && (abbr[i] >= '0' && abbr[i] <= '9') {
			num += string(abbr[i])
			i++
		}
		iNum, _ := strconv.Atoi(num)
		if iNum == 0 {
			return false
		}
		j += iNum
	}

	if i != len(abbr) || j != len(word) {
		return false
	}

	return true
}
