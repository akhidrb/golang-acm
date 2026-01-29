package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrack(&res, n, 0, 0, "")
	return res
}

func backtrack(res *[]string, n, i, j int, str string) {
	if i == j && j == n {
		*res = append(*res, str)
	}

	if i < n {
		str += "("
		backtrack(res, n, i+1, j, str)
		str = str[:len(str)-1]
	}
	if j < i {
		str += ")"
		backtrack(res, n, i, j+1, str)
	}
}
