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
		result := poisonousPlants(p)
		assert.Equal(t, int32(50000), result)
	})

	t.Run("2", func(t *testing.T) {
		p := readData("input_2")
		go func() {
			for {
				time.Sleep(500 * time.Millisecond)
				assert.Fail(t, "time limit exceeded")
				panic(errors.New("time limit exceeded"))
			}
		}()
		result := poisonousPlants(p)
		assert.Equal(t, int32(2), result)
	})
}

func readData(fileName string) []int32 {
	f, err := os.Open(fileName)
	checkError(err)
	reader := bufio.NewReaderSize(f, 16*1024*1024)
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var p []int32

	for i := 0; i < int(n); i++ {
		pItemTemp, err := strconv.ParseInt(pTemp[i], 10, 64)
		checkError(err)
		pItem := int32(pItemTemp)
		p = append(p, pItem)
	}
	return p
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
