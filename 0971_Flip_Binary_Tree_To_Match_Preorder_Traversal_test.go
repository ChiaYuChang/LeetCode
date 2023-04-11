package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFlipMatchVoyage(t *testing.T) {
	type testCase struct {
		name   string
		gen    func() *leetcode.Q0971TreeNode
		voyage []int
		answer []int
	}

	tcs := []testCase{
		{
			name: "Error at root",
			gen: func() *leetcode.Q0971TreeNode {
				root := &leetcode.Q0971TreeNode{Val: 1}
				root.Left = &leetcode.Q0971TreeNode{Val: 2}
				return root
			},
			voyage: []int{2, 1},
			answer: []int{-1},
		},
		{
			name: "One flip",
			gen: func() *leetcode.Q0971TreeNode {
				root := &leetcode.Q0971TreeNode{Val: 1}
				root.Left = &leetcode.Q0971TreeNode{Val: 2}
				root.Right = &leetcode.Q0971TreeNode{Val: 3}
				return root
			},
			voyage: []int{1, 3, 2},
			answer: []int{1},
		},
		{
			name: "No flip",
			gen: func() *leetcode.Q0971TreeNode {
				root := &leetcode.Q0971TreeNode{Val: 1}
				root.Left = &leetcode.Q0971TreeNode{Val: 2}
				root.Right = &leetcode.Q0971TreeNode{Val: 3}
				return root
			},
			voyage: []int{1, 2, 3},
			answer: []int{},
		},
		{
			name: "Not OK",
			gen: func() *leetcode.Q0971TreeNode {
				root := &leetcode.Q0971TreeNode{Val: 1}
				root.Left = &leetcode.Q0971TreeNode{Val: 2}
				root.Left.Left = &leetcode.Q0971TreeNode{Val: 5}
				root.Left.Right = &leetcode.Q0971TreeNode{Val: 4}

				root.Right = &leetcode.Q0971TreeNode{Val: 3}
				root.Right.Right = &leetcode.Q0971TreeNode{Val: 3}
				return root
			},
			voyage: []int{1, 2, 4, 5, 6, 3},
			answer: []int{-1},
		},
	}

	q := leetcode.Q0971{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				require.ElementsMatch(t, tc.answer, q.FlipMatchVoyage(tc.gen(), tc.voyage))
			},
		)
	}
}
