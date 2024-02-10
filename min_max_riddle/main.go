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
	n := len(arr)
	lowList := make([][]int64, n)
	for i := 0; i < n; i++ {
		lowList[i] = make([]int64, n)
	}
	for i := 0; i < n; i++ {
		least := arr[i]
		k := 0
		for j := i; j < n; j++ {
			if arr[j] <= least {
				least = arr[j]
			}
			lowList[k][i] = least
			k++
		}
	}
	highList := make([]int64, n)
	for i := 0; i < n; i++ {
		highest := lowList[i][0]
		for j := 0; j < n-i; j++ {
			if lowList[i][j] > highest {
				highest = lowList[i][j]
			}
		}
		highList[i] = highest
	}
	return highList
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
