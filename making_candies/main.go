package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumPasses' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. LONG_INTEGER m
 *  2. LONG_INTEGER w
 *  3. LONG_INTEGER p
 *  4. LONG_INTEGER n
 */

func minimumPasses(m int64, w int64, p int64, n int64) int64 {
	total := m * w
	passes := 1
	if total >= n {
		return int64(passes)
	}
	for total+(m*w) < n {
		pp := total / p
		total = total % p
		if pp%2 != 0 {
			if w < m {
				w++
			} else {
				m++
			}
			pp--
		}
		pp /= 2
		w += pp
		m += pp
		total += m * w
		passes++
	}
	return int64(passes + 1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("out")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	m, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)

	w, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)

	p, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)

	n, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)

	result := minimumPasses(m, w, p, n)

	fmt.Fprintf(writer, "%d\n", result)

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
