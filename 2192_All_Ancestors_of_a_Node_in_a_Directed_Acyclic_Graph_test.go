package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ2192(t *testing.T) {
	tcs := []struct {
		n      int
		edges  [][]int
		answer [][]int
	}{
		{
			n: 8,
			edges: [][]int{
				{0, 3},
				{0, 4},
				{1, 3},
				{2, 4},
				{2, 7},
				{3, 5},
				{3, 6},
				{3, 7},
				{4, 6},
			},
			answer: [][]int{
				{},
				{},
				{},
				{0, 1},
				{0, 2},
				{0, 1, 3},
				{0, 1, 2, 3, 4},
				{0, 1, 2, 3},
			},
		},
		{
			n: 5,
			edges: [][]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{0, 4},
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
			answer: [][]int{
				{},
				{0},
				{0, 1},
				{0, 1, 2},
				{0, 1, 2, 3},
			},
		},
	}

	q := leetcode.Q2192{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				result := q.GetAncestors(tc.n, tc.edges)
				require.Equal(t, len(tc.answer), len(result))
				for i := range result {
					require.Equal(t, tc.answer[i], result[i])
				}
			},
		)
	}
}
