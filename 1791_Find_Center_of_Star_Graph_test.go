package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ1791(t *testing.T) {
	tcs := []struct {
		edges  [][]int
		answer int
	}{
		{
			edges: [][]int{
				{1, 2}, {2, 3}, {4, 2},
			},
			answer: 2,
		},
		{
			edges: [][]int{
				{1, 2}, {5, 1}, {1, 3}, {1, 4},
			},
			answer: 1,
		},
	}

	q := leetcode.Q1791{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case_%d_nodes_star", len(tc.edges)),
			func(t *testing.T) {
				require.Equal(
					t, tc.answer, q.FindCenter(tc.edges),
				)
			},
		)
	}
}
