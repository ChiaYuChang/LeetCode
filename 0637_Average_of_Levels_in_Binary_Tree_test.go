package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestAverageOfLevels(t *testing.T) {
	NewNode := func(val int) *leetcode.Q0637TreeNode {
		return &leetcode.Q0637TreeNode{Val: val}
	}

	type testCase struct {
		genFunc func() (root *leetcode.Q0637TreeNode)
		answer  []float64
	}

	tcs := []testCase{
		{
			genFunc: func() (root *leetcode.Q0637TreeNode) {
				root = NewNode(3)
				root.Left = NewNode(9)
				root.Right = NewNode(20)
				root.Right.Right = NewNode(7)
				root.Right.Left = NewNode(15)
				return
			},
			answer: []float64{3.0, 14.5, 11.0},
		},
		{
			genFunc: func() (root *leetcode.Q0637TreeNode) {
				root = NewNode(3)
				root.Left = NewNode(9)
				root.Right = NewNode(20)
				root.Left.Left = NewNode(15)
				root.Left.Right = NewNode(7)
				return
			},
			answer: []float64{3.0, 14.5, 11.0},
		},
	}

	q := leetcode.Q0637{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()

				require.Equal(t, tc.answer, q.AverageOfLevels(tc.genFunc()))
			},
		)
	}
}
