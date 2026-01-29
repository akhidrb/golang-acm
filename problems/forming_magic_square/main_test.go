package main

import (
	"bufio"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		p := readData("input_1")
		go func() {
			for {
				time.Sleep(500 * time.Millisecond)
				assert.Fail(t, "time limit exceeded")
				panic(errors.New("time limit exceeded"))
			}
		}()
		result := formingMagicSquare(p)
		assert.Equal(t, int32(7), result)
	})
}

func readData(fileName string) [][]int32 {
	f, err := os.Open(fileName)
	checkError(err)
	reader := bufio.NewReaderSize(f, 16*1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}
	return s
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
