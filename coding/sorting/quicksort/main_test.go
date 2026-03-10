package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
		quicksort(nums)
		assert.Equal(t, []int{2, 3, 5, 7, 9, 10, 18, 101}, nums)
	})

}
