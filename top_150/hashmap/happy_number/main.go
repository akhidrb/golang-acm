package main

func isHappy(n int) bool {
	i := 0
	for n != 1 && i < 100 {
		sum := 0
		for n > 0 {
			dig := n % 10
			sum += dig * dig
			n /= 10
		}
		n = sum
		i++
	}
	return n == 1
}
