package main

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int64{6, 3, 5, 1, 12}
		res := riddle(nums)
		assert.Equal(t, []int64{12, 3, 3, 1, 1}, res)
	})
	t.Run("2", func(t *testing.T) {
		nums := []int64{2, 6, 1, 12}
		res := riddle(nums)
		assert.Equal(t, []int64{12, 2, 1, 1}, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := getData("input")
		exp := getData("output")
		res := riddle(nums)
		assert.Equal(t, exp, res)
	})

}

func getData(fileName string) []int64 {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	nums := make([]int64, 0)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return nil
		}
		nums = append(nums, int64(num))
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return nums
}
