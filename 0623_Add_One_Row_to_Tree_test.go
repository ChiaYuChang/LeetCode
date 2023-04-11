package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestAddOneRowToTree(t *testing.T) {
	var PreOrder func(root *leetcode.Q0623TreeNode, order *[]int)
	var PostOrder func(root *leetcode.Q0623TreeNode, order *[]int)

	PreOrder = func(root *leetcode.Q0623TreeNode, order *[]int) {
		if root == nil {
			return
		}
		(*order) = append((*order), root.Val)
		PreOrder(root.Left, order)
		PreOrder(root.Right, order)
	}

	PostOrder = func(root *leetcode.Q0623TreeNode, order *[]int) {
		if root == nil {
			return
		}
		PostOrder(root.Left, order)
		PostOrder(root.Right, order)
		(*order) = append((*order), root.Val)
	}

	newNode := func(val int) *leetcode.Q0623TreeNode {
		return &leetcode.Q0623TreeNode{Val: val}
	}

	type testCase struct {
		name      string
		genFun    func() (root *leetcode.Q0623TreeNode)
		value     int
		depth     int
		preOrder  []int
		postOrder []int
	}

	tcs := []testCase{
		{
			name: "No null node before target depth",
			genFun: func() (root *leetcode.Q0623TreeNode) {
				root = newNode(4)
				root.Left = newNode(2)
				root.Left.Left = newNode(3)
				root.Left.Right = newNode(1)

				root.Right = newNode(6)
				root.Right.Left = newNode(5)
				return root
			},
			value:     1,
			depth:     2,
			preOrder:  []int{4, 1, 2, 3, 1, 1, 6, 5},
			postOrder: []int{3, 1, 2, 1, 5, 6, 1, 4},
		},
		{
			name: "Null node before target depth",
			genFun: func() (root *leetcode.Q0623TreeNode) {
				root = newNode(4)
				root.Left = newNode(2)
				root.Left.Left = newNode(3)
				root.Left.Right = newNode(1)

				return root
			},
			value:     1,
			depth:     3,
			preOrder:  []int{4, 2, 1, 3, 1, 1},
			postOrder: []int{3, 1, 1, 1, 2, 4},
		},
		{
			name: "Depth = 1",
			genFun: func() (root *leetcode.Q0623TreeNode) {
				root = newNode(2)
				root.Left = newNode(3)
				root.Right = newNode(4)

				return root
			},
			value:     1,
			depth:     1,
			preOrder:  []int{1, 2, 3, 4},
			postOrder: []int{3, 4, 2, 1},
		},
	}

	q := leetcode.Q0623{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				tree := q.AddOneRow(tc.genFun(), tc.value, tc.depth)
				postOrder := []int{}
				PostOrder(tree, &postOrder)
				preOrder := []int{}
				PreOrder(tree, &preOrder)
				require.Equal(t, tc.postOrder, postOrder)
				require.Equal(t, tc.preOrder, preOrder)
			},
		)
	}
}
