package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		costs := []int{1, 2, 3, 4}
		n := 4
		res := climbStairs(n, costs)
		assert.Equal(t, 13, res)
	})

	t.Run("2", func(t *testing.T) {
		costs := []int{5, 1, 6, 2}
		n := 4
		res := climbStairs(n, costs)
		assert.Equal(t, 11, res)
	})

}
