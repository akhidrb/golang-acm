package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Main(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		}
		reorderList(&nums)
		expected := ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}
		assert.Equal(t, expected, nums)
	})

	t.Run("2", func(t *testing.T) {
		nums := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
		reorderList(&nums)
		expected := ListNode{Val: 1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}}
		assert.Equal(t, expected, nums)
	})

}
