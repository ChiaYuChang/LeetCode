package leetcode_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestTreeNodeDeepCopy(t *testing.T) {
	type testCase struct {
		SrcGenFun func() *leetcode.Q0894TreeNode
		CheckFun  func(t *testing.T, dst, src *leetcode.Q0894TreeNode)
	}

	tcs := []testCase{
		{
			SrcGenFun: func() *leetcode.Q0894TreeNode {
				src := &leetcode.Q0894TreeNode{Val: 1}
				return src
			},
			CheckFun: func(t *testing.T, dst, src *leetcode.Q0894TreeNode) {
				require.Equal(t, src.Val, dst.Val)
				require.Equal(t, src.Left, dst.Left)
				require.Equal(t, src.Right, dst.Right)

				src.Val = 2
				require.NotEqual(t, src.Val, dst.Val)
			},
		},
		{
			SrcGenFun: func() *leetcode.Q0894TreeNode {
				src := &leetcode.Q0894TreeNode{Val: 1}
				src.Left = &leetcode.Q0894TreeNode{Val: 2}
				src.Right = &leetcode.Q0894TreeNode{Val: 3}
				return src
			},
			CheckFun: func(t *testing.T, dst, src *leetcode.Q0894TreeNode) {
				require.Equal(t, src.Val, dst.Val)
				require.NotNil(t, dst.Left)
				require.NotNil(t, dst.Right)
				require.Equal(t, src.Left.Val, dst.Left.Val)
				require.Equal(t, src.Right.Val, dst.Right.Val)
			},
		},
		{
			SrcGenFun: func() *leetcode.Q0894TreeNode {
				src := &leetcode.Q0894TreeNode{Val: 1}
				src.Left = &leetcode.Q0894TreeNode{Val: 2}
				src.Left.Left = &leetcode.Q0894TreeNode{Val: 4}
				src.Left.Right = &leetcode.Q0894TreeNode{Val: 5}
				src.Right = &leetcode.Q0894TreeNode{Val: 3}
				src.Right.Right = &leetcode.Q0894TreeNode{Val: 6}
				return src
			},
			CheckFun: func(t *testing.T, dst, src *leetcode.Q0894TreeNode) {
				require.Equal(t, src.Val, dst.Val)
				require.NotNil(t, dst.Left)
				require.NotNil(t, dst.Right)
				require.Equal(t, src.Left.Val, dst.Left.Val)
				require.Equal(t, src.Right.Val, dst.Right.Val)
				require.NotNil(t, dst.Left.Left)
				require.NotNil(t, dst.Left.Right)
				require.Equal(t, src.Left.Left.Val, dst.Left.Left.Val)
				require.Equal(t, src.Left.Right.Val, dst.Left.Right.Val)
				require.NotNil(t, dst.Right.Right)
				require.Equal(t, src.Right.Right.Val, dst.Right.Right.Val)
			},
		},
	}

	q := leetcode.Q0984{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				src := tc.SrcGenFun()
				dst := q.TreeNodeDeepCopy(src)
				tc.CheckFun(t, src, dst)
			},
		)
	}
}

func TestAllPossibleFBT(t *testing.T) {
	InOrder := func(node *leetcode.Q0894TreeNode) string {
		order := []string{}

		var inOrder func(node *leetcode.Q0894TreeNode)
		inOrder = func(node *leetcode.Q0894TreeNode) {
			if node == nil {
				order = append(order, ".")
				return
			}
			order = append(order, strconv.Itoa(node.Val))
			inOrder(node.Left)
			inOrder(node.Right)
		}

		inOrder(node)
		return strings.Join(order, "")
	}

	PreOrder := func(node *leetcode.Q0894TreeNode) string {
		order := []string{}

		var preOrder func(node *leetcode.Q0894TreeNode)
		preOrder = func(node *leetcode.Q0894TreeNode) {
			if node == nil {
				order = append(order, ".")
				return
			}
			preOrder(node.Left)
			order = append(order, strconv.Itoa(node.Val))
			preOrder(node.Right)
		}

		preOrder(node)
		return strings.Join(order, "")
	}

	PostOrder := func(node *leetcode.Q0894TreeNode) string {
		order := []string{}

		var postOrder func(node *leetcode.Q0894TreeNode)
		postOrder = func(node *leetcode.Q0894TreeNode) {
			if node == nil {
				order = append(order, ".")
				return
			}
			postOrder(node.Left)
			postOrder(node.Right)
			order = append(order, strconv.Itoa(node.Val))
		}

		postOrder(node)
		return strings.Join(order, "")
	}

	type testCase struct {
		n          int
		inorder    []string
		preorder   []string
		postoreder []string
	}

	tcs := []testCase{
		{
			n:          1,
			inorder:    []string{"0.."},
			preorder:   []string{".0."},
			postoreder: []string{"..0"},
		},
		{
			n:          3,
			inorder:    []string{"00..0.."},
			preorder:   []string{".0.0.0."},
			postoreder: []string{"..0..00"},
		},
		{
			n: 5,
			inorder: []string{
				"00..00..0..", "000..0..0..",
			},
			preorder: []string{
				".0.0.0.0.0.", ".0.0.0.0.0.",
			},
			postoreder: []string{
				"..0..0..000", "..0..00..00",
			},
		},
		{
			n: 7,
			inorder: []string{
				"00..00..00..0..", "00..000..0..0..",
				"000..0..00..0..", "000..00..0..0..",
				"0000..0..0..0..",
			},
			preorder: []string{
				".0.0.0.0.0.0.0.", ".0.0.0.0.0.0.0.",
				".0.0.0.0.0.0.0.", ".0.0.0.0.0.0.0.",
				".0.0.0.0.0.0.0.",
			},
			postoreder: []string{
				"..0..0..0..0000", "..0..0..00..000",
				"..0..00..0..000", "..0..0..000..00",
				"..0..00..00..00",
			},
		},
	}

	q := leetcode.NewQ0984()
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-n=%d w/ cache", i+1, tc.n),
			func(t *testing.T) {
				trees := q.AllPossibleFBT(tc.n)
				inorder := make([]string, len(trees))
				preorder := make([]string, len(trees))
				postorder := make([]string, len(trees))
				for i, tree := range trees {
					inorder[i] = InOrder(tree)
					preorder[i] = PreOrder(tree)
					postorder[i] = PostOrder(tree)
				}
				require.ElementsMatch(t, tc.inorder, inorder)
				require.ElementsMatch(t, tc.postoreder, postorder)
				require.ElementsMatch(t, tc.preorder, preorder)
			},
		)
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-n=%d w/o cache", i+1, tc.n),
			func(t *testing.T) {
				t.Parallel()
				trees := q.AllPossibleFBT2(tc.n)
				inorder := make([]string, len(trees))
				preorder := make([]string, len(trees))
				postorder := make([]string, len(trees))
				for i, tree := range trees {
					inorder[i] = InOrder(tree)
					preorder[i] = PreOrder(tree)
					postorder[i] = PostOrder(tree)
				}
				require.ElementsMatch(t, tc.inorder, inorder)
				require.ElementsMatch(t, tc.postoreder, postorder)
				require.ElementsMatch(t, tc.preorder, preorder)
			},
		)
	}
}
