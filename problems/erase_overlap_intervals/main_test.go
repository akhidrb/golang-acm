package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		intervals := [][]int{
			{1, 2},
			{2, 3},
			{3, 4},
			{1, 3},
		}
		res := eraseOverlapIntervals(intervals)
		assert.Equal(t, 1, res)
	})

	t.Run("2", func(t *testing.T) {
		intervals := [][]int{
			{1, 11},
			{2, 12},
			{11, 22},
			{1, 100},
		}
		res := eraseOverlapIntervals(intervals)
		assert.Equal(t, 2, res)
	})

}
