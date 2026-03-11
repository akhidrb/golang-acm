package main

func myPow(x float64, n int) float64 {
	isNeg := false
	if n < 0 {
		isNeg = true
		n *= -1
	}
	res := powHelper(x, n)
	if isNeg {
		res = 1 / res
	}
	return res
}

func powHelper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := powHelper(x, n/2)
	res := half * half
	if n%2 == 1 {
		res *= x
	}
	return res
}
