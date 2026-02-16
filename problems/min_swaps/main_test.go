package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := "111000" // 101010 , 010101
		res := minSwaps(s)
		assert.Equal(t, 1, res)
	})

	t.Run("2", func(t *testing.T) {
		s := "010" // 010, 101
		res := minSwaps(s)
		assert.Equal(t, 0, res)
	})

	t.Run("3", func(t *testing.T) {
		s := "1110" // 1010, 0101
		res := minSwaps(s)
		assert.Equal(t, -1, res)
	})

	t.Run("4", func(t *testing.T) {
		s := "1100" // 1010
		res := minSwaps(s)
		assert.Equal(t, 1, res)
	})

	t.Run("5", func(t *testing.T) {
		s := "1100101100" // 1010101010, 0101010101
		res := minSwaps(s)
		assert.Equal(t, 2, res)
	})

	t.Run("6", func(t *testing.T) {
		s := "100" // 101, 010
		res := minSwaps(s)
		assert.Equal(t, 1, res)
	})
}

// 010, 101, 01, 10
