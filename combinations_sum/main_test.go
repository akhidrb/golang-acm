package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{2, 3, 6, 7}
		res := combinationSum(nums, 7)
		exp := [][]int{{2, 2, 3}, {7}}
		fmt.Println(res)
		fmt.Println(exp)
		if ok := checkArr(exp, res); !ok {
			assert.True(t, ok)
		}
		if ok := checkArr(res, exp); !ok {
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
