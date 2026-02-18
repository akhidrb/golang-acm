package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		strs := []string{"flower", "flow", "flight"}
		res := longestCommonPrefix(strs)
		assert.Equal(t, "fl", res)
	})

	t.Run("2", func(t *testing.T) {
		strs := []string{"dog", "racecar", "car"}
		res := longestCommonPrefix(strs)
		assert.Equal(t, "", res)
	})

	t.Run("3", func(t *testing.T) {
		strs := []string{"a"}
		res := longestCommonPrefix(strs)
		assert.Equal(t, "a", res)
	})

	t.Run("4", func(t *testing.T) {
		strs := []string{"ab", "a"}
		res := longestCommonPrefix(strs)
		assert.Equal(t, "a", res)
	})

	t.Run("5", func(t *testing.T) {
		strs := []string{"flower", "flower", "flower", "flower"}
		res := longestCommonPrefix(strs)
		assert.Equal(t, "flower", res)
	})

}
