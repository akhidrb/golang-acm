package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RandomizedTest(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		actions := []string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"}
		nums := [][]int{{}, {1}, {2}, {2}, {}, {1}, {2}, {}}
		res := handleActions(actions, nums)
		expectedRes := []interface{}{nil, true, false, true, 2, true, false, 2}
		assert.Equal(t, expectedRes, res)
	})
}

func handleActions(actions []string, nums [][]int) []interface{} {
	resList := make([]interface{}, 0)
	var obj RandomizedSet
	for i, action := range actions {
		var res interface{}
		switch action {
		case "RandomizedSet":
			obj = Constructor()
			break
		case "insert":
			res = obj.Insert(nums[i][0])
			break
		case "remove":
			res = obj.Remove(nums[i][0])
			break
		case "getRandom":
			res = obj.GetRandom()
			break
		}
		resList = append(resList, res)
	}
	return resList
}
