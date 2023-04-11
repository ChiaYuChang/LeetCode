package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestTree2Str(t *testing.T) {
	type testCase struct {
		fun func() *leetcode.Q0606TreeNode
		ans string
	}

	tcs := []testCase{
		{
			fun: func() *leetcode.Q0606TreeNode {
				root := &leetcode.Q0606TreeNode{Val: 1}
				root.Left = &leetcode.Q0606TreeNode{Val: 2}
				root.Right = &leetcode.Q0606TreeNode{Val: 3}
				root.Left.Left = &leetcode.Q0606TreeNode{Val: 4}
				return root
			},
			ans: "1(2(4))(3)",
		},
		{
			fun: func() *leetcode.Q0606TreeNode {
				root := &leetcode.Q0606TreeNode{Val: 1}
				root.Left = &leetcode.Q0606TreeNode{Val: 2}
				root.Right = &leetcode.Q0606TreeNode{Val: 3}
				root.Left.Right = &leetcode.Q0606TreeNode{Val: 4}
				return root
			},
			ans: "1(2()(4))(3)",
		},
		{
			fun: func() *leetcode.Q0606TreeNode {
				root := &leetcode.Q0606TreeNode{Val: 1}
				root.Left = &leetcode.Q0606TreeNode{Val: 2}
				root.Right = &leetcode.Q0606TreeNode{Val: 3}
				root.Left.Right = &leetcode.Q0606TreeNode{Val: 4}
				root.Right.Right = &leetcode.Q0606TreeNode{Val: 5}
				return root
			},
			ans: "1(2()(4))(3()(5))",
		},
		{
			fun: func() *leetcode.Q0606TreeNode {
				root := &leetcode.Q0606TreeNode{Val: 1}
				return root
			},
			ans: "1",
		},
		{
			fun: func() *leetcode.Q0606TreeNode {
				return nil
			},
			ans: "",
		},
	}

	q := leetcode.Q0606{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.Tree2Str(tc.fun()))
			},
		)
	}
}
