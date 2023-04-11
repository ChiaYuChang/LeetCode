package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLevelOrderBottom(t *testing.T) {
	NewNode := func(val int) *leetcode.Q0107TreeNode {
		return &leetcode.Q0107TreeNode{Val: val}
	}

	type testCase struct {
		genFunc func() (root *leetcode.Q0107TreeNode)
		answer  [][]int
	}

	tcs := []testCase{
		{
			genFunc: func() (root *leetcode.Q0107TreeNode) {
				root = NewNode(3)
				root.Left = NewNode(9)
				root.Right = NewNode(20)
				root.Right.Left = NewNode(15)
				root.Right.Right = NewNode(7)
				return
			},
			answer: [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			genFunc: func() (root *leetcode.Q0107TreeNode) {
				root = NewNode(1)
				return
			},
			answer: [][]int{{1}},
		},
		{
			genFunc: func() (root *leetcode.Q0107TreeNode) {
				return
			},
			answer: [][]int{},
		},
	}

	q := leetcode.Q0107{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.answer, q.LevelOrderBottom(tc.genFunc()))
			},
		)
	}
}
