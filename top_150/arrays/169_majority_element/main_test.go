package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{3, 2, 3}
		k := majorityElement(nums)
		assert.Equal(t, 3, k)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{2, 2, 1, 1, 1, 2, 2}
		k := majorityElement(nums)
		assert.Equal(t, 2, k)
	})

}
