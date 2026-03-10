package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		nums := [][]int{{0, 30}, {5, 10}, {15, 20}}
		rooms := minMeetingRooms(nums)
		assert.Equal(t, 2, rooms)
	})

	t.Run("2", func(t *testing.T) {
		nums := [][]int{{7, 10}, {2, 4}}
		rooms := minMeetingRooms(nums)
		assert.Equal(t, 1, rooms)
	})

}
