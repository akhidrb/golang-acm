package main

import (
	"fmt"
	"strconv"
)

func Solution(S string) string {
	numList := make([]int, 10)
	for _, val := range S {
		numList[int(val)-'0']++
	}
	first, last := "", ""
	oneRem, oneRemStr := true, ""
	for i := 9; i >= 0; i-- {
		reps := numList[i] / 2
		for j := 0; j < reps; j++ {
			if i != 0 || len(first) > 0 {
				first += strconv.Itoa(i)
				last = strconv.Itoa(i) + last
			}
		}
		if rem := numList[i] % 2; rem > 0 && oneRem {
			oneRem = false
			oneRemStr = strconv.Itoa(i)
		}
	}
	str := first + oneRemStr + last
	if str == "" {
		return "0"
	}
	return str
}

func main() {
	fmt.Println(Solution("39878"))
	fmt.Println(Solution("009009"))
	fmt.Println(Solution("0090009"))
	fmt.Println(Solution("00900"))
	fmt.Println(Solution("0000"))
	fmt.Println(Solution("54321"))
	fmt.Println(Solution("8199"))
	fmt.Println(Solution("83238"))
	fmt.Println(Solution("44"))
	fmt.Println(Solution("7"))
}
