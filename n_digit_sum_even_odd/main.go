package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	globeList []int
)

func main() {
	n := 3
	myList := loop(n)
	fmt.Println(myList)
	globeList = make([]int, 0)
	recursion("", n, 0, 0, 0)
	fmt.Println(globeList)
}

func loop(n int) []int {
	myList := make([]int, 0)
	minNumber := int(math.Pow(10, float64(n-1)))
	maxNumber := int(math.Pow(10, float64(n)) - 1)
	for i := minNumber; i <= maxNumber; i++ {
		total := map[int]int{
			0: 0,
			1: 0,
		}
		s2 := strconv.Itoa(i)
		for pos, v := range s2 {
			numPos, _ := strconv.Atoi(string(v))
			total[pos%2] += numPos
		}
		if total[0] == total[1] {
			myList = append(myList, i)
		}
	}
	return myList
}

func recursion(s string, n, odd, even, flag int) {
	if n > 0 {
		d := '0'
		if s == "" {
			d = '1'
		}
		for ; d <= '9'; d++ {
			if flag == 1 {
				recursion(s+string(d), n-1, odd+int(d-'0'), even, 0)
			} else {
				recursion(s+string(d), n-1, odd, even+int(d-'0'), 1)
			}
		}
	}
	if n == 0 && even == odd {
		num, _ := strconv.Atoi(s)
		globeList = append(globeList, num)
	}
}
