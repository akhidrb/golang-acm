package main

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}

	// Map DNA char -> 2-bit code
	code := [256]int{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}

	const L = 10
	const mask = (1 << (2 * L)) - 1

	seen := make(map[int]int, len(s))
	res := make([]string, 0)

	x := 0
	for i := 0; i < len(s); i++ {
		x = ((x << 2) | code[s[i]]) & mask

		if i >= L-1 {
			seen[x]++
			if seen[x] == 2 {
				res = append(res, s[i-L+1:i+1])
			}
		}
	}

	return res
}
