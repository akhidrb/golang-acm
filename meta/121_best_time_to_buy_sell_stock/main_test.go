package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{7, 1, 5, 3, 6, 4}
		res := maxProfit(nums)
		assert.Equal(t, 5, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{7, 6, 4, 3, 1}
		res := maxProfit(nums)
		assert.Equal(t, 0, res)
	})

}
