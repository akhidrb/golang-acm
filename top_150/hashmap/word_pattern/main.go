package main

import "strings"

func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	if len(pattern) != len(arr) {
		return false
	}
	dict := make(map[byte]string)
	dict2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		if val, ok := dict[p]; ok {
			if val != arr[i] {
				return false
			}
		} else {
			dict[p] = arr[i]
		}

		if val, ok := dict2[arr[i]]; ok {
			if val != p {
				return false
			}
		} else {
			dict2[arr[i]] = p
		}

	}
	return true
}
