package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev := 0
	num := x
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return rev == num
}
