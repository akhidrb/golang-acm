package main

func minSwaps(s string) int {
	valid1, valid2 := constructValidStrs(len(s))
	count1, count2 := 0, 0
	zeros1, zeros2 := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' && s[i] != valid1[i] {
			count1++
		}
		if s[i] == '0' && s[i] != valid1[i] {
			zeros1++
		}
		if s[i] == '1' && s[i] != valid2[i] {
			count2++
		}
		if s[i] == '0' && s[i] != valid2[i] {
			zeros2++
		}
	}
	if count1 != zeros1 && count2 != zeros2 {
		return -1
	} else if count1 != zeros1 {
		return count2
	} else if count2 != zeros2 {
		return count1
	}
	count1 = min(count1, count2)
	return count1
}

func constructValidStrs(n int) ([]byte, []byte) {
	valid1 := make([]byte, n)
	valid2 := make([]byte, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			valid1[i] = '1'
			valid2[i] = '0'
		} else {
			valid1[i] = '0'
			valid2[i] = '1'
		}
	}
	return valid1, valid2
}
