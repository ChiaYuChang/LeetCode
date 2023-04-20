package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMaxLevelSum(t *testing.T) {
	NewNode := func(val int) *leetcode.Q1161TreeNode {
		return &leetcode.Q1161TreeNode{Val: val}
	}

	type testCase struct {
		genFunc func() (root *leetcode.Q1161TreeNode)
		answer  int
	}

	tcs := []testCase{
		{
			genFunc: func() (root *leetcode.Q1161TreeNode) {
				root = NewNode(1)
				root.Left = NewNode(7)
				root.Right = NewNode(0)
				root.Left.Left = NewNode(7)
				root.Left.Right = NewNode(-8)
				return
			},
			answer: 2,
		},
		{
			genFunc: func() (root *leetcode.Q1161TreeNode) {
				root = NewNode(989)
				root.Right = NewNode(10250)
				root.Right.Left = NewNode(98693)
				root.Right.Right = NewNode(-89388)
				root.Right.Right.Right = NewNode(-32127)
				return
			},
			answer: 2,
		},
	}

	q := leetcode.Q1161{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.answer, q.MaxLevelSum(tc.genFunc()))
			},
		)
	}
}
