package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := minimumPasses(1, 1, 6, 45)
		exp := int64(16)
		assert.Equal(t, exp, res)
	})

	t.Run("2", func(t *testing.T) {
		res := minimumPasses(1, 2, 1, 60)
		exp := int64(4)
		assert.Equal(t, exp, res)
	})

	t.Run("3", func(t *testing.T) {
		res := minimumPasses(3, 1, 2, 12)
		exp := int64(3)
		assert.Equal(t, exp, res)
	})
}
