package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLevelOrder(t *testing.T) {
	NewNode := func(val int) *leetcode.Node {
		return &leetcode.Node{Val: val}
	}

	NewChildren := func(vals ...int) []*leetcode.Node {
		ns := make([]*leetcode.Node, len(vals))
		for i, val := range vals {
			ns[i] = NewNode(val)
		}
		return ns
	}

	type testCase struct {
		genFunc func() (root *leetcode.Node)
		answer  [][]int
	}

	tcs := []testCase{
		{
			genFunc: func() (root *leetcode.Node) {
				root = NewNode(1)
				root.Children = NewChildren(3, 2, 4)
				root.Children[0].Children = NewChildren(5, 6)
				return
			},
			answer: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			genFunc: func() (root *leetcode.Node) {
				root = NewNode(1)
				root.Children = NewChildren(2, 3, 4, 5)

				root.Children[1].Children = NewChildren(6, 7)
				root.Children[1].Children[1].Children = append(root.Children[1].Children[1].Children, NewNode(11))
				root.Children[1].Children[1].Children[0].Children = append(root.Children[1].Children[1].Children[0].Children, NewNode(14))

				root.Children[2].Children = append(root.Children[2].Children, NewNode(8))
				root.Children[2].Children[0].Children = append(root.Children[2].Children[0].Children, NewNode(12))

				root.Children[3].Children = NewChildren(9, 10)
				root.Children[3].Children[0].Children = append(root.Children[3].Children[0].Children, NewNode(13))
				return
			},
			answer: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
	}

	q := leetcode.Q0429{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.answer, q.LevelOrder(tc.genFunc()))
			},
		)
	}
}
