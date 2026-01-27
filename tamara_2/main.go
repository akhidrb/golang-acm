package main

import (
	"bytes"
	"fmt"
)

func largestPalindromic(S string) string {
	numList := make([]int, 10)
	for _, val := range S {
		numList[int(val)-'0']++
	}
	leftHalf := make([]byte, 0)
	oneRem := ""
	for i := 9; i >= 0; i-- {
		reps := numList[i] / 2
		for j := 0; j < reps; j++ {
			if i != 0 || len(leftHalf) > 0 {
				leftHalf = append(leftHalf, byte(i)+'0')
			}
		}
		if rem := numList[i] % 2; rem > 0 && oneRem == "" {
			oneRem = string(byte(i) + '0')
		}
	}
	rightHalf := reverseStr(leftHalf)
	str := string(leftHalf) + oneRem + rightHalf
	if str == "" {
		return "0"
	}
	return str
}

func reverseStr(left []byte) string {
	var sb bytes.Buffer
	n := len(left) - 1
	for i := n; i >= 0; i-- {
		sb.WriteString(string(left[i]))
	}
	return sb.String()
}

func main() {
	fmt.Println(largestPalindromic("39878"))
	fmt.Println(largestPalindromic("009009"))
	fmt.Println(largestPalindromic("0090009"))
	fmt.Println(largestPalindromic("00900"))
	fmt.Println(largestPalindromic("0000"))
	fmt.Println(largestPalindromic("54321"))
	fmt.Println(largestPalindromic("8199"))
	fmt.Println(largestPalindromic("83238"))
	fmt.Println(largestPalindromic("44"))
	fmt.Println(largestPalindromic("7"))
}
