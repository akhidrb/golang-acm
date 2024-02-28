package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3}
		res := permute(nums)
		exp := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
		if ok := checkArr(exp, res); !ok {
			fmt.Println(exp)
			fmt.Println(res)
			assert.True(t, ok)
		}
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0, 1}
		res := permute(nums)
		exp := [][]int{{0, 1}, {1, 0}}
		if ok := checkArr(exp, res); !ok {
			fmt.Println(exp)
			fmt.Println(res)
			assert.True(t, ok)
		}
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{1}
		res := permute(nums)
		exp := [][]int{{1}}
		if ok := checkArr(exp, res); !ok {
			fmt.Println(exp)
			fmt.Println(res)
			assert.True(t, ok)
		}
	})
}

func checkArr(exp, act [][]int) bool {
	if len(exp) != len(act) {
		return false
	}
	m := make(map[string]bool)
	for _, val := range exp {
		key := toString(val)
		m[key] = true
	}
	for _, val := range act {
		key := toString(val)
		if _, ok := m[key]; !ok {
			return false
		}
	}
	return true
}

func toString(a []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ""), "[]")
}
