package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		valid := merge(nums)
		exp := [][]int{{1, 6}, {8, 10}, {15, 18}}
		assert.Equal(t, exp, valid)
	})

	t.Run("2", func(t *testing.T) {
		nums := [][]int{{1, 4}, {4, 5}}
		valid := merge(nums)
		exp := [][]int{{1, 5}}
		assert.Equal(t, exp, valid)
	})

	t.Run("3", func(t *testing.T) {
		nums := [][]int{{4, 7}, {1, 4}}
		valid := merge(nums)
		exp := [][]int{{1, 7}}
		assert.Equal(t, exp, valid)
	})

}
