package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestKthLargestLevelSum(t *testing.T) {
	NewNode := func(val int) *leetcode.Q1161TreeNode {
		return &leetcode.Q1161TreeNode{Val: val}
	}

	type testCase struct {
		genFunc func() (root *leetcode.Q1161TreeNode)
		k       int
		answer  int64
	}

	tcs := []testCase{
		{
			genFunc: func() (root *leetcode.Q1161TreeNode) {
				root = NewNode(5)
				root.Left = NewNode(8)
				root.Left.Left = NewNode(2)
				root.Left.Left.Left = NewNode(4)
				root.Left.Left.Right = NewNode(6)
				root.Left.Right = NewNode(1)
				root.Right = NewNode(9)
				root.Right.Left = NewNode(3)
				root.Right.Right = NewNode(7)
				return
			},
			k:      2,
			answer: 13,
		},
		{
			genFunc: func() (root *leetcode.Q1161TreeNode) {
				root = NewNode(1)
				root.Left = NewNode(2)
				root.Left.Left = NewNode(3)
				return
			},
			k:      1,
			answer: 3,
		},
		{
			genFunc: func() (root *leetcode.Q1161TreeNode) {
				root = NewNode(5)
				root.Left = NewNode(8)
				root.Left.Left = NewNode(2)
				root.Left.Right = NewNode(1)
				root.Right = NewNode(9)
				root.Right.Left = NewNode(3)
				root.Right.Right = NewNode(7)
				return
			},
			k:      4,
			answer: -1,
		},
	}

	q := leetcode.Q2583{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.answer, q.KthLargestLevelSum(tc.genFunc(), tc.k))
			},
		)
	}
}
