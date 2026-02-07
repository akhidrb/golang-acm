package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		coins := []int{1, 2, 5}
		amount := 11
		res := coinChange(coins, amount)
		assert.Equal(t, 3, res)
	})

	t.Run("1", func(t *testing.T) {
		coins := []int{2}
		amount := 3
		res := coinChange(coins, amount)
		assert.Equal(t, -1, res)
	})

	t.Run("1", func(t *testing.T) {
		coins := []int{1}
		amount := 0
		res := coinChange(coins, amount)
		assert.Equal(t, 0, res)
	})

}
