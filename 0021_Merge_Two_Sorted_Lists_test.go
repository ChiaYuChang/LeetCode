package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMergeTwoList(t *testing.T) {
	NewNode := func(v int) *leetcode.Q0021ListNode {
		return &leetcode.Q0021ListNode{Val: v}
	}

	NewLinkedList := func(nums []int) *leetcode.Q0021ListNode {
		if len(nums) == 0 {
			return nil
		}
		root := NewNode(nums[0])
		curr := root
		for i := 1; i < len(nums); i++ {
			curr.Next = NewNode(nums[i])
			curr = curr.Next
		}
		return root
	}

	var requireListEqual func(t *testing.T, l1, l2 *leetcode.Q0021ListNode)
	requireListEqual = func(t *testing.T, l1, l2 *leetcode.Q0021ListNode) {
		l1IsNil, l2IsNil := l1 == nil, l2 == nil

		// l1IsNil and l2IsNil should be equal
		require.Equal(t, l1IsNil, l2IsNil)

		if l1IsNil == false && l2IsNil == false {
			// check value
			require.Equal(t, l1.Val, l2.Val)
		} else {
			// both nil
			return
		}
		requireListEqual(t, l1.Next, l2.Next)
	}

	type testCase struct {
		name string
		l1   *leetcode.Q0021ListNode
		l2   *leetcode.Q0021ListNode
		ans  *leetcode.Q0021ListNode
	}

	tcs := []testCase{
		{
			name: "General case",
			l1:   NewLinkedList([]int{1, 2, 4}),
			l2:   NewLinkedList([]int{1, 3, 4}),
			ans:  NewLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "Two empty list",
			l1:   NewLinkedList([]int{}),
			l2:   NewLinkedList([]int{}),
			ans:  NewLinkedList([]int{}),
		},
		{
			name: "Single empty list",
			l1:   NewLinkedList([]int{}),
			l2:   NewLinkedList([]int{0}),
			ans:  NewLinkedList([]int{0}),
		},
	}

	q := leetcode.Q0021{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				requireListEqual(t, tc.ans, q.MergeTwoLists(tc.l1, tc.l2))
			},
		)
	}
}
