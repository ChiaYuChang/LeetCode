package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindDiagonalOrder(t *testing.T) {
	type testCase struct {
		matrix [][]int
		order  []int
	}

	tcs := []testCase{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			order:  []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			matrix: [][]int{{1, 2}, {3, 4}},
			order:  []int{1, 2, 3, 4},
		},
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			order:  []int{1, 2, 5, 9, 6, 3, 4, 7, 10, 11, 8, 12},
		},
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			order:  []int{1, 2, 4, 7, 5, 3, 6, 8, 10, 11, 9, 12},
		},
		{
			matrix: [][]int{{1}},
			order:  []int{1},
		},
		{
			matrix: [][]int{{1, 2}},
			order:  []int{1, 2},
		},
		{
			matrix: [][]int{{1}, {2}},
			order:  []int{1, 2},
		},
	}

	q := leetcode.Q0498{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(
					t, tc.order,
					q.FindDiagonalOrder(tc.matrix),
				)
			},
		)
	}
}
