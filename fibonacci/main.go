package main

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibMemo(n int) int {
	memo := make(map[int]int)
	return memoRecur(n, memo)
}

func memoRecur(n int, memo map[int]int) int {
	// check in memo, if found, retrieve and return right away
	if val, ok := memo[n]; ok {
		return val
	}

	if n == 0 || n == 1 {
		return n
	}

	res := memoRecur(n-1, memo) + memoRecur(n-2, memo)

	// save result to memo before returning
	memo[n] = res
	return res
}
