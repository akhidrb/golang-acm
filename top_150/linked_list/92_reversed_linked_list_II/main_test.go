package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		left, right := 2, 4
		head := createLinkedListFromSlice(arr)
		res := reverseBetween(head, left, right)
		expList := []int{1, 4, 3, 2, 5}
		exp := createLinkedListFromSlice(expList)
		assert.Equal(t, *exp, *res)
	})

	t.Run("2", func(t *testing.T) {
		arr := []int{5}
		left, right := 1, 1
		head := createLinkedListFromSlice(arr)
		res := reverseBetween(head, left, right)
		assert.Equal(t, ListNode{Val: 5}, *res)
	})

	t.Run("3", func(t *testing.T) {
		arr := []int{3, 5}
		left, right := 1, 2
		head := createLinkedListFromSlice(arr)
		res := reverseBetween(head, left, right)
		expList := []int{5, 3}
		exp := createLinkedListFromSlice(expList)
		assert.Equal(t, *exp, *res)
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
