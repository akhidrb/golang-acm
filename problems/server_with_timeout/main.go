package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
 * Complete the 'SyncHash' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. int32 steps
 *  2. chan Result resultChan0
 *  3. chan Result resultChan1
 */

func SyncHash(steps int, resultChan0 chan Result, resultChan1 chan Result) (map[string]int32, error) {
	m := make(map[string]int32)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs error

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < steps; i++ {
			timeBegin := time.Now()
			r1 := <-resultChan0
			timeEnd := time.Now()
			mu.Lock()
			if _, ok := m[r1.Hash]; ok {
				m[r1.Hash] += r1.Number
			} else {
				m[r1.Hash] = r1.Number
			}
			mu.Unlock()
			if timeEnd.After(timeBegin.Add(maxDelay)) {
				errs = timeoutErr
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < steps; i++ {
			timeBegin := time.Now()
			r2 := <-resultChan1
			timeEnd := time.Now()
			mu.Lock()
			if _, ok := m[r2.Hash]; ok {
				m[r2.Hash] += r2.Number
			} else {
				m[r2.Hash] = r2.Number
			}
			mu.Unlock()
			if timeEnd.After(timeBegin.Add(maxDelay)) {
				errs = timeoutErr
				break
			}
		}
	}()
	wg.Wait()

	if errs != nil {
		return map[string]int32{}, errs
	}

	return m, nil
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("out")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	hashesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var hashes []string

	for i := 0; i < int(hashesCount); i++ {
		hashesItem := readLine(reader)
		hashes = append(hashes, hashesItem)
	}

	numbers0Count, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var numbers0 []int32

	for i := 0; i < int(numbers0Count); i++ {
		numbers0ItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		numbers0Item := int32(numbers0ItemTemp)
		numbers0 = append(numbers0, numbers0Item)
	}

	numbers1Count, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var numbers1 []int32

	for i := 0; i < int(numbers1Count); i++ {
		numbers1ItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		numbers1Item := int32(numbers1ItemTemp)
		numbers1 = append(numbers1, numbers1Item)
	}

	thresholdTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	threshold := int32(thresholdTemp)

	server0, server1 := makeServer(), makeServer()
	resultChan0 := server0(hashes, numbers0, threshold)
	resultChan1 := server1(hashes, numbers1, threshold)
	result, err := SyncHash(len(hashes), resultChan0, resultChan1)
	if err == nil {
		for _, hash := range hashes {
			fmt.Fprintf(writer, hash)
			fmt.Fprintf(writer, " ")
			fmt.Fprintf(writer, fmt.Sprintf("%d", result[hash]))
			fmt.Fprintf(writer, "\n")
		}
	} else {
		fmt.Fprintf(writer, err.Error())
	}

	writer.Flush()
}

var timeoutErr = errors.New("Timeout error")
var maxDelay = 100 * time.Millisecond

type Result struct {
	Hash   string
	Number int32
}

func makeServer() func([]string, []int32, int32) chan Result {
	return func(hashes []string, numbers []int32, threshold int32) chan Result {
		resultChan := make(chan Result)
		go func() {
			indexes := []int{}
			for i := 0; i < len(hashes); i++ {
				indexes = append(indexes, i)
			}
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(indexes), func(i, j int) { indexes[i], indexes[j] = indexes[j], indexes[i] })
			for _, i := range indexes {
				if numbers[i] > threshold {
					time.Sleep(125 * time.Millisecond)
				} else {
					time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
				}
				resultChan <- Result{
					Hash:   hashes[i],
					Number: numbers[i],
				}
			}
		}()
		return resultChan
	}
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
