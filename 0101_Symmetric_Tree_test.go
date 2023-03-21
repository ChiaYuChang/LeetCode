package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestIsSymmetricTree(t *testing.T) {
	NewNode := func(val int) *leetcode.Q0101TreeNode {
		return &leetcode.Q0101TreeNode{Val: val}
	}

	r1 := NewNode(1)
	r1.Left = NewNode(2)
	r1.Left.Left = NewNode(3)
	r1.Left.Right = NewNode(4)
	r1.Right = NewNode(2)
	r1.Right.Left = NewNode(4)
	r1.Right.Right = NewNode(3)

	r2 := NewNode(1)
	r2.Left = NewNode(2)
	r2.Right = NewNode(2)
	r2.Right = NewNode(3)
	r2.Right = NewNode(3)

	r3 := NewNode(1)

	r4 := NewNode(1)
	r4.Left = NewNode(2)

	type testCase struct {
		name        string
		root        *leetcode.Q0101TreeNode
		isSymmetric bool
	}

	tcs := []testCase{
		{
			name:        "Symmetric Tree",
			root:        r1,
			isSymmetric: true,
		},
		{
			name:        "Asymmetric Tree",
			root:        r2,
			isSymmetric: false,
		},
		{
			name:        "Single node tree",
			root:        r3,
			isSymmetric: true,
		},
		{
			name:        "Nil in one branch",
			root:        r4,
			isSymmetric: false,
		},
	}

	q := leetcode.Q0101{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				require.Equal(
					t, tc.isSymmetric,
					q.IsSymmetric(tc.root),
				)
			},
		)
	}
}
