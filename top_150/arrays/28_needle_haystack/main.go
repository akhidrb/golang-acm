package main

func strStr(haystack string, needle string) int {
	i := 0
	n := len(needle)
	for i+n <= len(haystack) {
		str := haystack[i : i+n]
		if str == needle {
			return i
		}
		i++
	}
	return -1
}
