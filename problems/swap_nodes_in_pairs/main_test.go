package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LongestConsecutiveSequence(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		listNode := createLinkedListFromSlice(nums)
		res := swapPairs(listNode)
		expected := createLinkedListFromSlice([]int{2, 1, 4, 3})
		assert.Equal(t, expected, res)
	})

	t.Run("2", func(t *testing.T) {
		res := swapPairs(nil)
		var listNode *ListNode
		assert.Equal(t, listNode, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{1}
		listNode := createLinkedListFromSlice(nums)
		res := swapPairs(listNode)
		expected := createLinkedListFromSlice([]int{1})
		assert.Equal(t, expected, res)
	})

	t.Run("4", func(t *testing.T) {
		nums := []int{2, 5, 3, 4, 6, 2, 2}
		listNode := createLinkedListFromSlice(nums)
		res := swapPairs(listNode)
		expected := createLinkedListFromSlice([]int{5, 2, 4, 3, 2, 6, 2})
		assert.Equal(t, expected, res)
	})

}

func createLinkedListFromSlice(nums []int) *ListNode {
	if nums == nil {
		return nil
	}
	l := &ListNode{}
	p := l
	for i, _ := range nums {
		p.Val = nums[i]
		if i < len(nums)-1 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}
	return l
}
