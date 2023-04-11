package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestIsBalanced(t *testing.T) {
	type testCase struct {
		name    string
		genFunc func() *leetcode.Q0110TreeNode
		answer  bool
	}

	NewNode := func(n int) *leetcode.Q0110TreeNode {
		return &leetcode.Q0110TreeNode{Val: n}
	}

	tcs := []testCase{
		{
			name: "Balanced tree",
			genFunc: func() *leetcode.Q0110TreeNode {
				root := NewNode(3)
				root.Left = NewNode(9)
				root.Right = NewNode(10)
				root.Right.Left = NewNode(15)
				root.Right.Right = NewNode(17)
				return root
			},
			answer: true,
		},
		{
			name: "Unbalanced tree",
			genFunc: func() *leetcode.Q0110TreeNode {
				root := NewNode(1)
				root.Left = NewNode(2)
				root.Left.Left = NewNode(3)
				root.Left.Left.Left = NewNode(4)
				root.Left.Left.Right = NewNode(4)
				root.Left.Right = NewNode(3)
				root.Right = NewNode(2)
				return root
			},
			answer: false,
		},
		{
			name: "Null",
			genFunc: func() *leetcode.Q0110TreeNode {
				return nil
			},
			answer: true,
		},
	}

	q := leetcode.Q0110{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.answer, q.IsBalanced(tc.genFunc()))
			},
		)
	}
}
