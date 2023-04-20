package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestWidthOfBinaryTree(t *testing.T) {
	NewNode := func(val int) *leetcode.Q0662TreeNode {
		return &leetcode.Q0662TreeNode{Val: val}
	}

	type testCase struct {
		Name          string
		TreeBuildFunc func() *leetcode.Q0662TreeNode
		MaxWidth      int
	}

	tcs := []testCase{
		{
			Name: "Simple Case I",
			TreeBuildFunc: func() *leetcode.Q0662TreeNode {
				root := NewNode(1)
				root.Left = NewNode(3)
				root.Right = NewNode(2)

				root.Left.Left = NewNode(5)
				root.Left.Right = NewNode(3)

				root.Right = NewNode(2)
				root.Right.Right = NewNode(9)
				return root
			},
			MaxWidth: 4,
		},
		{
			Name: "Simple Case II",
			TreeBuildFunc: func() *leetcode.Q0662TreeNode {
				root := NewNode(1)
				root.Left = NewNode(3)
				root.Right = NewNode(2)

				root.Left.Left = NewNode(5)
				root.Left.Left.Left = NewNode(6)

				root.Right.Right = NewNode(9)
				root.Right.Right.Left = NewNode(7)
				return root
			},
			MaxWidth: 7,
		},
		{
			Name: "Single node at the bottom",
			TreeBuildFunc: func() *leetcode.Q0662TreeNode {
				root := NewNode(1)
				root.Left = NewNode(3)
				root.Right = NewNode(2)
				root.Left.Left = NewNode(5)
				return root
			},
			MaxWidth: 2,
		},
		{
			Name: "Only root node",
			TreeBuildFunc: func() *leetcode.Q0662TreeNode {
				root := NewNode(1)
				return root
			},
			MaxWidth: 1,
		},
		{
			Name: "Extreme deep tree",
			TreeBuildFunc: func() *leetcode.Q0662TreeNode {
				root := NewNode(0)
				root.Left = NewNode(0)
				root.Right = NewNode(0)
				l, r := root.Left, root.Right
				for i := 0; i < 100; i++ {
					l.Right = NewNode(0)
					r.Left = NewNode(0)
					l, r = l.Right, r.Left
				}
				return root
			},
			MaxWidth: 2,
		},
		{
			Name: "Extreme deep tree with multiple leaves",
			TreeBuildFunc: func() *leetcode.Q0662TreeNode {
				n := NewNode(0)
				n.Left = NewNode(1)
				n.Right = NewNode(2)

				n.Left.Left = NewNode(3)
				n.Left.Right = NewNode(4)

				n.Right.Left = NewNode(5)
				n.Right.Right = NewNode(6)

				root := NewNode(0)
				curr := root
				for i := 0; i < 100; i++ {
					curr.Right = NewNode(0)
					curr = curr.Right
				}
				curr.Right = n
				return root
			},
			MaxWidth: 4,
		},
	}

	q := leetcode.Q0662{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.Name),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.MaxWidth, q.WidthOfBinaryTree(tc.TreeBuildFunc()))
			},
		)
	}
}
