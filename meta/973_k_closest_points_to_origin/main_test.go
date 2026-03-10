package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := [][]int{{1, 3}, {-2, 2}}
		valid := kClosest(nums, 1)
		exp := [][]int{{-2, 2}}
		assert.ElementsMatch(t, exp, valid)
	})

	t.Run("2", func(t *testing.T) {
		nums := [][]int{{3, 3}, {5, -1}, {-2, 4}}
		valid := kClosest(nums, 2)
		exp := [][]int{{3, 3}, {-2, 4}}
		assert.ElementsMatch(t, exp, valid)
	})

	t.Run("3", func(t *testing.T) {
		nums := [][]int{{0, 1}, {1, 0}}
		valid := kClosest(nums, 2)
		exp := [][]int{{0, 1}, {1, 0}}
		assert.ElementsMatch(t, exp, valid)
	})

	t.Run("4", func(t *testing.T) {
		nums := [][]int{{1, 3}, {-2, 2}}
		valid := kClosestSorting(nums, 1)
		exp := [][]int{{-2, 2}}
		assert.ElementsMatch(t, exp, valid)
	})

	t.Run("5", func(t *testing.T) {
		nums := [][]int{{3, 3}, {5, -1}, {-2, 4}}
		valid := kClosestSorting(nums, 2)
		exp := [][]int{{3, 3}, {-2, 4}}
		assert.ElementsMatch(t, exp, valid)
	})

	t.Run("6", func(t *testing.T) {
		nums := [][]int{{0, 1}, {1, 0}}
		valid := kClosestSorting(nums, 2)
		exp := [][]int{{0, 1}, {1, 0}}
		assert.ElementsMatch(t, exp, valid)
	})

}
