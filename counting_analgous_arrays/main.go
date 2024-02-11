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
 * Complete the 'countAnalogousArrays' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY consecutiveDifference
 *  2. INTEGER lowerBound
 *  3. INTEGER upperBound
 */

func countAnalogousArrays(consecutiveDifference []int32, lowerBound int32, upperBound int32) int32 {
	low, high := lowerBound, upperBound
	total := 0
	for low <= high {
		sum := low
		higher := false
		for _, val := range consecutiveDifference {
			sum -= val
			if sum > upperBound {
				higher = true
				break
			} else if sum < lowerBound {
				higher = true
				break
			}
		}
		if !higher {
			total++
		}
		low++
	}
	return int32(total)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("out")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	consecutiveDifferenceCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var consecutiveDifference []int32

	for i := 0; i < int(consecutiveDifferenceCount); i++ {
		consecutiveDifferenceItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		consecutiveDifferenceItem := int32(consecutiveDifferenceItemTemp)
		consecutiveDifference = append(consecutiveDifference, consecutiveDifferenceItem)
	}

	lowerBoundTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	lowerBound := int32(lowerBoundTemp)

	upperBoundTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	upperBound := int32(upperBoundTemp)

	result := countAnalogousArrays(consecutiveDifference, lowerBound, upperBound)

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
