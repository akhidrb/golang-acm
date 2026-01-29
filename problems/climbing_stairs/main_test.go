package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := climbStairs(2)
		assert.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		res := climbStairs(3)
		assert.Equal(t, 3, res)
	})

	t.Run("3", func(t *testing.T) {
		res := climbStairs(4)
		assert.Equal(t, 5, res)
	})

	t.Run("4", func(t *testing.T) {
		res := climbStairs(44)
		assert.Equal(t, 1134903170, res)
	})

}
