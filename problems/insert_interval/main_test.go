package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Main(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		interval := [][]int{{1, 3}, {6, 9}}
		newInterval := []int{2, 5}
		actual := insert(interval, newInterval)
		expected := [][]int{{1, 5}, {6, 9}}
		assert.Equal(t, expected, actual)
	})

}
