package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		n := 2
		res := grayCode(n)
		assert.Equal(t, []int{0, 1, 3, 2}, res)
	})

	t.Run("2", func(t *testing.T) {
		n := 1
		res := grayCode(n)
		assert.Equal(t, []int{0, 1}, res)
	})
}
