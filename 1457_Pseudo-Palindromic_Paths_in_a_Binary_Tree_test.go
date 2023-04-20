package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestPseudoPalindromicPaths(t *testing.T) {
	NewNode := func(v int) *leetcode.Q1457TreeNode {
		return &leetcode.Q1457TreeNode{Val: v}
	}

	type testCase struct {
		TreeBuildFunc func() *leetcode.Q1457TreeNode
		answer        int
	}

	tcs := []testCase{
		{
			TreeBuildFunc: func() *leetcode.Q1457TreeNode {
				root := NewNode(2)
				root.Left = NewNode(3)
				root.Right = NewNode(1)

				root.Left.Left = NewNode(3)
				root.Left.Right = NewNode(1)
				root.Right.Right = NewNode(1)

				return root
			},
			answer: 2,
		},
		{
			TreeBuildFunc: func() *leetcode.Q1457TreeNode {
				root := NewNode(2)
				root.Left = NewNode(1)
				root.Right = NewNode(1)

				root.Left.Left = NewNode(1)
				root.Left.Right = NewNode(3)
				root.Left.Right.Right = NewNode(1)

				return root
			},
			answer: 1,
		},
		{
			TreeBuildFunc: func() *leetcode.Q1457TreeNode {
				root := NewNode(9)
				return root
			},
			answer: 1,
		},
	}

	q := leetcode.Q1457{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(
					t, tc.answer,
					q.PseudoPalindromicPaths(tc.TreeBuildFunc()),
				)
			},
		)
	}
}
