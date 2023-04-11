package leetcode_test

import (
	"sort"
	"testing"

	"gitlab.com/gjerry134679/leetcode"
)

func TestMergeKLists(t *testing.T) {
	NewNode := func(n int) *leetcode.Q0023ListNode {
		return &leetcode.Q0023ListNode{Val: n}
	}

	SliceToListNode := func(nums []int) *leetcode.Q0023ListNode {
		if len(nums) == 0 {
			return nil
		}

		if !sort.IsSorted(sort.IntSlice(nums)) {
			sort.Sort(sort.IntSlice(nums))
		}

		head := NewNode(nums[0])
		curr := head
		for i := 1; i < len(nums); i++ {
			curr.Next = NewNode(nums[i])
			curr = curr.Next
		}
		return head
	}

	ListNodeToSlice := func(ln *leetcode.Q0023ListNode) []int {
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
		lists  [][]int
		answer []int
	}

	tcs := []testCase{
		{
			lists:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			answer: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			lists:  [][]int{},
			answer: []int{},
		},
		{
			lists:  [][]int{{}},
			answer: []int{},
		},
	}

	q := leetcode.Q0023{}
	for i := range tcs {
		tc := tcs[i]
		lists := make([]*leetcode.Q0023ListNode, len(tc.lists))
		for i, l := range tc.lists {
			lists[i] = SliceToListNode(l)
		}
		t.Logf("require:\n\t- want: %v\n\t- get : %v",
			tc.answer,
			ListNodeToSlice(q.MergeKLists(lists)),
		)

	}
}
