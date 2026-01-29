package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{3, 2, 2, 3}
		val := 3
		res := removeElement(nums, val)
		k := 2
		assert.Equal(t, k, res)
		numsExp := []int{2, 2}
		nums = nums[:k]
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		for i := 0; i < k; i++ {
			assert.Equal(t, numsExp[i], nums[i])
		}
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
		val := 2
		res := removeElement(nums, val)
		k := 5
		assert.Equal(t, k, res)
		numsExp := []int{0, 0, 1, 3, 4}
		nums = nums[:k]
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		for i := 0; i < k; i++ {
			assert.Equal(t, numsExp[i], nums[i])
		}
	})

}
