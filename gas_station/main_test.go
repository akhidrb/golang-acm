package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GasStation(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		cost := []int{3, 4, 5, 1, 2}
		res := canCompleteCircuit(nums, cost)
		assert.Equal(t, 3, res)
	})
	t.Run("2", func(t *testing.T) {
		nums := []int{2, 3, 4}
		cost := []int{3, 4, 3}
		res := canCompleteCircuit(nums, cost)
		assert.Equal(t, -1, res)
	})
	t.Run("3", func(t *testing.T) {
		nums := []int{3, 1, 1}
		cost := []int{1, 2, 2}
		res := canCompleteCircuit(nums, cost)
		assert.Equal(t, 0, res)
	})
}
