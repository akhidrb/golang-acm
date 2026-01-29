package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var ans [26]int
	for i, val := range t {
		ans[val-'a']++
		ans[s[i]-'a']--
	}
	for _, val := range ans {
		if val != 0 {
			return false
		}
	}
	return true
}
