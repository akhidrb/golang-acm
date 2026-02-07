package main

func isIsomorphic(s string, t string) bool {
	check1 := check(s, t)
	check2 := check(t, s)
	return check1 && check2
}

func check(s string, t string) bool {
	cache := make(map[byte]byte)
	for i := range s {
		char1 := s[i]
		if val, ok := cache[char1]; ok {
			if val != t[i] {
				return false
			}
		} else {
			cache[char1] = t[i]
		}
	}
	return true
}
