package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMergeKList(t *testing.T) {
	NewNode := func(v int) *leetcode.Q0023ListNode {
		return &leetcode.Q0023ListNode{Val: v}
	}

	NewLinkedList := func(nums []int) *leetcode.Q0023ListNode {
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

	var requireListEqual func(t *testing.T, l1, l2 *leetcode.Q0023ListNode)
	requireListEqual = func(t *testing.T, l1, l2 *leetcode.Q0023ListNode) {
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
		name  string
		lists []*leetcode.Q0023ListNode
		ans   *leetcode.Q0023ListNode
	}

	tcs := []testCase{
		{
			lists: []*leetcode.Q0023ListNode{
				NewLinkedList([]int{1, 4, 5}),
				NewLinkedList([]int{1, 3, 4}),
				NewLinkedList([]int{2, 6}),
			},
			ans: NewLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
	}

	q := leetcode.Q0023{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				requireListEqual(t, tc.ans, q.MergeLists(tc.lists...))
			},
		)
	}
}
