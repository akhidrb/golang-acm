package main

func minMovesToMakePalindrome(s string) int {
	p1, p2 := 0, len(s)-1
	arr := []byte(s)
	count := 0
	for p1 < p2 {
		j := p2
		for j > p1 && arr[p1] != arr[j] {
			j--
		}

		if j == p1 {
			arr[p1], arr[p1+1] = arr[p1+1], arr[p1]
			count++
			continue
		}

		for j < p2 {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j++
			count++
		}
		p1++
		p2--
	}
	return count
}
