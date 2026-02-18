package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		k := 2
		head := createLinkedListFromSlice(arr)
		res := rotateRight(head, k)
		expList := []int{4, 5, 1, 2, 3}
		exp := createLinkedListFromSlice(expList)
		assert.Equal(t, *exp, *res)
	})

	t.Run("2", func(t *testing.T) {
		arr := []int{0, 1, 2}
		k := 4
		head := createLinkedListFromSlice(arr)
		res := rotateRight(head, k)
		expList := []int{2, 0, 1}
		exp := createLinkedListFromSlice(expList)
		assert.Equal(t, *exp, *res)
	})

	t.Run("3", func(t *testing.T) {
		arr := []int{1, 2, 3}
		k := 2000000000
		head := createLinkedListFromSlice(arr)
		res := rotateRight(head, k)
		expList := []int{2, 0, 1}
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
