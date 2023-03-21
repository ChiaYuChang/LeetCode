package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func cyclicLiskListBuilder(vals []int, pos int) (*leetcode.Q0142ListNode, *leetcode.Q0142ListNode) {
	if len(vals) == 0 {
		return nil, nil
	}

	var branchnode *leetcode.Q0142ListNode
	root := &leetcode.Q0142ListNode{Val: vals[0]}
	curr := root

	for i := 1; i < len(vals); i++ {
		curr.Next = &leetcode.Q0142ListNode{Val: vals[i]}
		curr = curr.Next
		if i == pos {
			branchnode = curr
		}
	}

	switch pos {
	case -1:
		curr.Next = nil
	case 0:
		curr.Next = root
		branchnode = root
	default:
		curr.Next = branchnode
	}
	return root, branchnode
}

func TestDetectCycle(t *testing.T) {
	type testCase struct {
		name string
		vals []int
		pos  int
	}

	tcs := []testCase{
		{
			name: "General",
			vals: []int{3, 2, 0, -4},
			pos:  1,
		},
		{
			name: "Root",
			vals: []int{1, 2},
			pos:  0,
		},
		{
			name: "Single elemnt acyclic list",
			vals: []int{1},
			pos:  -1,
		},
		{
			name: "General acyclic list",
			vals: []int{1, 2},
			pos:  -1,
		},
		{
			name: "Null list",
			vals: []int{},
			pos:  -1,
		},
	}

	q := leetcode.Q0142{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				root, brch := cyclicLiskListBuilder(tc.vals, tc.pos)

				if brch == nil {
					require.Nil(t, q.DetectCycle(root))
				} else {
					require.Equal(t, brch, q.DetectCycle(root))
				}
			},
		)
	}
}
