package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		num := "112358"
		res := isAdditiveNumber(num)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		num := "199100199"
		res := isAdditiveNumber(num)
		assert.Equal(t, true, res)
	})

	t.Run("3", func(t *testing.T) {
		num := "199100198"
		res := isAdditiveNumber(num)
		assert.Equal(t, false, res)
	})

	t.Run("4", func(t *testing.T) {
		num := "1023"
		res := isAdditiveNumber(num)
		assert.Equal(t, false, res)
	})

	t.Run("5", func(t *testing.T) {
		num := "112"
		res := isAdditiveNumber(num)
		assert.Equal(t, true, res)
	})
}
