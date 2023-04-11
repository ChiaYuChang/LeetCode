package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestIsBipartite(t *testing.T) {
	type testCase struct {
		name  string
		graph [][]int
		ans   bool
	}

	tcs := []testCase{
		{
			name: "Case 1",
			graph: [][]int{
				{1, 2, 3},
				{0, 2},
				{0, 1, 3},
				{0, 2},
			},
			ans: false,
		},
		{
			name: "Case 2",
			graph: [][]int{
				{1, 3},
				{0, 2},
				{1, 3},
				{0, 2},
			},
			ans: true,
		},
		{
			name: "w/ multiple nodes with 0 out degree",
			graph: [][]int{
				{3},
				{7, 9},
				{},
				{0, 5},
				{},
				{3},
				{9},
				{1},
				{},
				{1, 6},
			},
			ans: true,
		},
		{
			name: "w/ one node with 0 out degree",
			graph: [][]int{
				{4},
				{},
				{4},
				{4},
				{0, 2, 3},
			},
			ans: true,
		},
		{
			name: "first node is 0 out degree node",
			graph: [][]int{
				{},
				{2, 4, 6},
				{1, 4, 8, 9},
				{7, 8},
				{1, 2, 8, 9},
				{6, 9},
				{1, 5, 7, 8, 9},
				{3, 6, 9},
				{2, 3, 4, 6, 9},
				{2, 4, 5, 6, 7, 8},
			},
			ans: false,
		},
	}

	q := leetcode.Q0785{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.IsBipartite(tc.graph))
			},
		)
	}
}
