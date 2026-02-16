package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		low, high := 100, 300
		res := sequentialDigits(low, high)
		assert.Equal(t, []int{123, 234}, res)
	})

	t.Run("2", func(t *testing.T) {
		low, high := 1000, 13000
		res := sequentialDigits(low, high)
		assert.Equal(t, []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}, res)
	})

	t.Run("3", func(t *testing.T) {
		low, high := 58, 155
		res := sequentialDigits(low, high)
		assert.Equal(t, []int{67, 78, 89, 123}, res)
	})

	t.Run("4", func(t *testing.T) {
		low, high := 8511, 23553
		res := sequentialDigits(low, high)
		assert.Equal(t, []int{12345, 23456}, res)
	})
}
