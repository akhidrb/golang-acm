package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		numCourses := 2
		prerequisites := [][]int{{1, 0}}
		res := canFinish(numCourses, prerequisites)
		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		numCourses := 2
		prerequisites := [][]int{{1, 0}, {0, 1}}
		res := canFinish(numCourses, prerequisites)
		assert.Equal(t, false, res)
	})

}
