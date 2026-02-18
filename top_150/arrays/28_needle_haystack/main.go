package main

func strStr(haystack string, needle string) int {
	n := len(needle)
	for i := 0; i+n <= len(haystack); i++ {
		str := haystack[i : i+n]
		if str == needle {
			return i
		}
	}
	return -1
}
