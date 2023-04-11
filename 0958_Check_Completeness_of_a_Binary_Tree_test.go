package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestIsCompleteTree(t *testing.T) {
	type testCase struct {
		name string
		gen  func() *leetcode.Q0958TreeNode
		ans  bool
	}

	tcs := []testCase{
		{
			name: "Root is nil",
			gen: func() *leetcode.Q0958TreeNode {
				return nil
			},
			ans: true,
		},
		{
			name: "Only root",
			gen: func() *leetcode.Q0958TreeNode {
				return &leetcode.Q0958TreeNode{Val: 1}
			},
			ans: true,
		},
		{
			name: "Perfect tree",
			gen: func() *leetcode.Q0958TreeNode {
				root := &leetcode.Q0958TreeNode{Val: 1}
				root.Left = &leetcode.Q0958TreeNode{Val: 2}
				root.Right = &leetcode.Q0958TreeNode{Val: 3}
				return root
			},
			ans: true,
		},

		{
			name: "Complete tree",
			gen: func() *leetcode.Q0958TreeNode {
				root := &leetcode.Q0958TreeNode{Val: 1}

				root.Left = &leetcode.Q0958TreeNode{Val: 2}
				root.Left.Left = &leetcode.Q0958TreeNode{Val: 4}
				root.Left.Right = &leetcode.Q0958TreeNode{Val: 5}

				root.Right = &leetcode.Q0958TreeNode{Val: 3}
				root.Right.Left = &leetcode.Q0958TreeNode{Val: 6}
				return root
			},
			ans: true,
		},
		{
			name: "Incomplete tree I",
			gen: func() *leetcode.Q0958TreeNode {
				root := &leetcode.Q0958TreeNode{Val: 1}

				root.Left = &leetcode.Q0958TreeNode{Val: 2}
				root.Left.Left = &leetcode.Q0958TreeNode{Val: 4}
				root.Left.Right = &leetcode.Q0958TreeNode{Val: 5}
				root.Left.Left.Left = &leetcode.Q0958TreeNode{Val: 7}

				root.Right = &leetcode.Q0958TreeNode{Val: 3}
				root.Right.Left = &leetcode.Q0958TreeNode{Val: 6}
				return root
			},
			ans: false,
		},
		{
			name: "Incomplete tree II",
			gen: func() *leetcode.Q0958TreeNode {
				root := &leetcode.Q0958TreeNode{Val: 1}

				root.Left = &leetcode.Q0958TreeNode{Val: 2}
				root.Left.Left = &leetcode.Q0958TreeNode{Val: 4}
				root.Left.Right = &leetcode.Q0958TreeNode{Val: 5}

				root.Right = &leetcode.Q0958TreeNode{Val: 3}
				root.Right.Right = &leetcode.Q0958TreeNode{Val: 6}
				return root
			},
			ans: false,
		},
	}

	q := leetcode.Q0958{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.IsCompleteTree(tc.gen()))
			},
		)
	}
}
