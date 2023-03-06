package leetcode_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSet(t *testing.T) {
	type testCase struct {
		name string
		i    []int
		ans  []int
	}

	tcs := []testCase{
		{
			name: "w/o repeated elements",
			i:    []int{0, 1, 2, 3},
			ans:  []int{0, 1, 2, 3},
		},
		{
			name: "w/ repeated elements",
			i:    []int{1, 0, 0, 1, 1, 1, 2, 3, 2, 3},
			ans:  []int{0, 1, 2, 3},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				s := leetcode.NewSet()
				require.True(t, s.IsEmpty())
				s.Add(tc.i...)
				require.False(t, s.IsEmpty())
				ans := s.ToList()
				sort.Sort(sort.IntSlice(ans))
				require.Equal(t, len(tc.ans), len(ans))
				for i := range tc.ans {
					require.Equal(t, tc.ans[i], ans[i])
				}

				for _, e := range tc.ans {
					s.RemoveOne(e)
				}
				require.True(t, s.IsEmpty())
			},
		)
	}
}

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
				{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9},
				{6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9},
				{2, 4, 5, 6, 7, 8},
			},
			ans: false,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t, tc.ans, leetcode.IsBipartite(tc.graph))
			},
		)
	}
}

func TestGraphPartition(t *testing.T) {
	graph := [][]int{
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
	}
	graphGroups := leetcode.GarphPartition(graph)
	for _, g := range graphGroups {
		t.Log(g)
	}
}
