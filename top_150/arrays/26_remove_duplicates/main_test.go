package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 1, 2}
		k := removeDuplicates(nums)
		assert.Equal(t, 2, k)
		expNums := []int{1, 2}
		nums = nums[:k]
		assert.Equal(t, expNums, nums)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		k := removeDuplicates(nums)
		assert.Equal(t, 5, k)
		expNums := []int{0, 1, 2, 3, 4}
		nums = nums[:k]
		assert.Equal(t, expNums, nums)
	})

}
