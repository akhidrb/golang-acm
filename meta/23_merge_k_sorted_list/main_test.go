package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		arr := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
		lists := make([]*ListNode, len(arr))
		for _, value := range arr {
			lists = append(lists, createLinkedListFromSlice(value))
		}
		res := mergeKLists(lists)
		expList := []int{1, 1, 2, 3, 4, 4, 5, 6}
		assert.Equal(t, createLinkedListFromSlice(expList), res)
	})

}

func createLinkedListFromSlice(nums []int) *ListNode {
	if nums == nil {
		return nil
	}
	l := &ListNode{}
	p := l
	for i := range nums {
		p.Val = nums[i]
		if i < len(nums)-1 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}
	return l
}
