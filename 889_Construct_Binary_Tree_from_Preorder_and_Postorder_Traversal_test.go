package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestPreAndPostOrder(t *testing.T) {
	root := &leetcode.Q889TreeNode{Val: 0}
	curn := root
	curn.Left = &leetcode.Q889TreeNode{Val: 1}
	curn.Right = &leetcode.Q889TreeNode{Val: 2}

	curn = root.Left
	curn.Left = &leetcode.Q889TreeNode{Val: 3}
	curn.Right = &leetcode.Q889TreeNode{Val: 4}

	curn = root.Right
	curn.Left = &leetcode.Q889TreeNode{Val: 5}

	curn = root.Left.Left
	curn.Left = &leetcode.Q889TreeNode{Val: 6}
	curn.Right = &leetcode.Q889TreeNode{Val: 7}

	curn = root.Right.Left
	curn.Right = &leetcode.Q889TreeNode{Val: 8}

	preorder := []int{0, 1, 3, 6, 7, 4, 2, 5, 8}
	postorder := []int{6, 7, 3, 4, 1, 8, 5, 2, 0}

	q := leetcode.Q889{}
	for i, n := range q.PreOrder(root) {
		require.Equal(t, preorder[i], n)
	}

	for i, n := range q.PostOrder(root) {
		require.Equal(t, postorder[i], n)
	}
}

func TestConstructFromPrePost(t *testing.T) {
	type testCase struct {
		preorder  []int
		postorder []int
	}

	tcs := []testCase{
		{
			preorder:  []int{1, 2, 4, 5, 3, 6, 7},
			postorder: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			preorder:  []int{0, 1, 3, 6, 7, 4, 2, 5, 8},
			postorder: []int{6, 7, 3, 4, 1, 8, 5, 2, 0},
		},
		{
			preorder:  []int{0, 1, 2, 3, 4, 5},
			postorder: []int{5, 4, 3, 2, 1, 0},
		},
		{
			preorder:  []int{1},
			postorder: []int{1},
		},
	}

	q := leetcode.Q889{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				tree := q.ConstructFromPrePost(tc.preorder, tc.postorder)
				require.Equal(t, tc.postorder, q.PostOrder(tree))
				require.Equal(t, tc.preorder, q.PreOrder(tree))
			},
		)
	}
}
