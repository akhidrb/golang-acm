package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		digits := []int{1, 2, 3}
		res := plusOne(digits)
		exp := []int{1, 2, 4}
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		digits := []int{4, 3, 2, 1}
		res := plusOne(digits)
		exp := []int{4, 3, 2, 2}
		assert.Equal(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		digits := []int{9}
		res := plusOne(digits)
		exp := []int{1, 0}
		assert.Equal(t, exp, res)
	})

	t.Run("4", func(t *testing.T) {
		digits := []int{1, 2, 9}
		res := plusOne(digits)
		exp := []int{1, 3, 0}
		assert.Equal(t, exp, res)
	})

	t.Run("5", func(t *testing.T) {
		digits := []int{9, 9}
		res := plusOne(digits)
		exp := []int{1, 0, 0}
		assert.Equal(t, exp, res)
	})

	t.Run("6", func(t *testing.T) {
		digits := []int{2, 3, 9, 9}
		res := plusOne(digits)
		exp := []int{2, 4, 0, 0}
		assert.Equal(t, exp, res)
	})
}
