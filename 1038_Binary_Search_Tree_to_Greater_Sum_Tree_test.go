package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ1038(t *testing.T) {
	q := leetcode.Q1038{}

	tcs := []struct {
		root   *leetcode.Q1038TreeNode
		answer string
	}{
		{
			root: &leetcode.Q1038TreeNode{
				4,
				&leetcode.Q1038TreeNode{
					1,
					&leetcode.Q1038TreeNode{0, nil, nil},
					&leetcode.Q1038TreeNode{2, nil, &leetcode.Q1038TreeNode{3, nil, nil}},
				},
				&leetcode.Q1038TreeNode{
					6,
					&leetcode.Q1038TreeNode{5, nil, nil},
					&leetcode.Q1038TreeNode{7, nil, &leetcode.Q1038TreeNode{8, nil, nil}},
				},
			},
			answer: "[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]",
		},
		{
			root: &leetcode.Q1038TreeNode{
				0,
				nil,
				&leetcode.Q1038TreeNode{1, nil, nil},
			},
			answer: "[1,null,1]",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case_%d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.answer, q.BSTToGst(tc.root).String())
			},
		)
	}

}
