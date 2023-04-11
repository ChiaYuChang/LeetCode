package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestRotateRight(t *testing.T) {
	NewNode := func(n int) *leetcode.Q0061ListNode {
		return &leetcode.Q0061ListNode{Val: n}
	}

	SliceToListNode := func(nums []int) *leetcode.Q0061ListNode {
		if len(nums) == 0 {
			return nil
		}
		head := NewNode(nums[0])
		curr := head
		for i := 1; i < len(nums); i++ {
			curr.Next = NewNode(nums[i])
			curr = curr.Next
		}
		return head
	}

	ListNodeToSlice := func(ln *leetcode.Q0061ListNode) []int {
		if ln == nil {
			return []int{}
		}
		slc := []int{}
		for curr := ln; curr != nil; curr = curr.Next {
			slc = append(slc, curr.Val)
		}
		return slc
	}

	type testCase struct {
		head []int
		k    int
		ans  []int
	}

	tcs := []testCase{
		{
			head: []int{1, 2, 3, 4, 5},
			k:    2,
			ans:  []int{4, 5, 1, 2, 3},
		},
		{
			head: []int{0, 1, 2},
			k:    4,
			ans:  []int{2, 0, 1},
		},
		{
			head: []int{},
			k:    0,
			ans:  []int{},
		},
		{
			head: []int{1},
			k:    0,
			ans:  []int{1},
		},
		{
			head: []int{1, 2},
			k:    0,
			ans:  []int{1, 2},
		},
		{
			head: []int{1, 2},
			k:    3,
			ans:  []int{2, 1},
		},
	}

	q := leetcode.Q0061{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans,
					ListNodeToSlice(q.RotateRight(SliceToListNode(tc.head), tc.k)))
			},
		)
	}
}
