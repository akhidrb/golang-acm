package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func riddle(arr []int64) []int64 {
	wMax := make(map[int64]int64)
	stk := make([][]int64, 0)
	var i int64
	n := int64(len(arr))
	arr = append(arr, 0)
	for i = 0; i < int64(len(arr)); i++ {
		idx := i
		for len(stk) > 0 && stk[len(stk)-1][0] > arr[i] {
			val, prevIdx := stk[len(stk)-1][0], stk[len(stk)-1][1]
			stk = stk[:len(stk)-1]
			temp1 := i - prevIdx + 1
			if w, ok := wMax[arr[i]]; ok {
				if w > temp1 {
					temp1 = w
				}
			}
			wMax[arr[i]] = temp1

			temp2 := i - prevIdx
			if w, ok := wMax[val]; ok {
				if w > temp2 {
					temp2 = w
				}
			}
			wMax[val] = temp2
			idx = prevIdx
		}
		stk = append(stk, [][]int64{{arr[i], idx}}...)
	}
	delete(wMax, 0)
	wInv := make(map[int64]int64)
	for key, value := range wMax {
		if w, ok := wInv[value]; ok {
			if w > key {
				key = w
			}
		}
		wInv[value] = key
	}
	ll := make([]int64, n)
	ll[0] = wInv[n]
	for i = n - 1; i >= 1; i-- {
		temp := ll[n-i-1]
		if w, ok := wInv[i]; ok {
			if w > temp {
				temp = w
			}
		}
		ll[n-i] = temp
	}
	var j int64
	for i, j = 0, n-1; i < j; i, j = i+1, j-1 {
		ll[i], ll[j] = ll[j], ll[i]
	}
	return ll
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	res := riddle(arr)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
