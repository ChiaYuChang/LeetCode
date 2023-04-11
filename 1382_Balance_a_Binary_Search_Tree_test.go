package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestBuildBalanceBST(t *testing.T) {
	type testCase struct {
		genFunc func() (root *leetcode.Q1382TreeNode)
		inOrder []int
	}

	NewNode := func(val int) *leetcode.Q1382TreeNode {
		return &leetcode.Q1382TreeNode{Val: val}
	}

	var InOrder func(node *leetcode.Q1382TreeNode, list *[]int)
	InOrder = func(node *leetcode.Q1382TreeNode, list *[]int) {
		if node == nil {
			return
		}
		InOrder(node.Left, list)
		(*list) = append((*list), node.Val)
		InOrder(node.Right, list)
	}

	IsBalance := func(root *leetcode.Q1382TreeNode) bool {
		var helper func(root *leetcode.Q1382TreeNode) int
		helper = func(root *leetcode.Q1382TreeNode) int {
			if root == nil {
				return 0
			}

			lhs := helper(root.Left)
			if lhs == -1 {
				return -1
			}

			rhs := helper(root.Right)
			if rhs == -1 {
				return -1
			}

			if rhs > lhs {
				rhs, lhs = lhs, rhs
			}

			if lhs-rhs > 1 {
				return -1
			}
			return lhs + 1
		}
		return helper(root) != -1
	}

	tcs := []testCase{
		{
			func() (root *leetcode.Q1382TreeNode) {
				root = NewNode(1)
				root.Right = NewNode(2)
				root.Right.Right = NewNode(3)
				root.Right.Right.Right = NewNode(4)
				return root
			},
			[]int{1, 2, 3, 4},
		},
		{
			func() (root *leetcode.Q1382TreeNode) {
				root = NewNode(2)
				root.Left = NewNode(1)
				root.Right = NewNode(3)
				return root
			},
			[]int{1, 2, 3},
		},
	}

	q := leetcode.Q1382{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				tree := q.BalanceBST(tc.genFunc())
				require.True(t, IsBalance(tree))
				order := []int{}
				InOrder(tree, &order)
				require.Equal(t, tc.inOrder, order)
			},
		)
	}
}
