package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LongestConsecutiveSequence(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		n := 2
		listNode := createLinkedListFromSlice(nums)
		res := removeNthFromEnd(listNode, n)
		expected := createLinkedListFromSlice([]int{1, 2, 3, 5})
		assert.Equal(t, expected, res)
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1}
		n := 1
		listNode := createLinkedListFromSlice(nums)
		res := removeNthFromEnd(listNode, n)
		expected := createLinkedListFromSlice(nil)
		assert.Equal(t, expected, res)
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{1, 2}
		n := 1
		listNode := createLinkedListFromSlice(nums)
		res := removeNthFromEnd(listNode, n)
		expected := createLinkedListFromSlice([]int{1})
		assert.Equal(t, expected, res)
	})
	t.Run("4", func(t *testing.T) {
		nums := []int{1, 2}
		n := 2
		listNode := createLinkedListFromSlice(nums)
		res := removeNthFromEnd(listNode, n)
		expected := createLinkedListFromSlice([]int{2})
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
