package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestNetworkBecomesIdle(t *testing.T) {
	type testCase struct {
		edges    [][]int
		patience []int
		answer   int
	}

	tcs := []testCase{
		{
			edges: [][]int{
				{0, 1},
				{1, 2},
			},
			patience: []int{
				0, 2, 1,
			},
			answer: 8,
		},
		{
			edges: [][]int{
				{0, 1},
				{0, 2},
				{1, 2},
			},
			patience: []int{
				0, 10, 10,
			},
			answer: 3,
		},
		{
			edges: [][]int{
				{3, 8},
				{4, 13},
				{0, 7},
				{0, 4},
				{1, 8},
				{14, 1},
				{7, 2},
				{13, 10},
				{9, 11},
				{12, 14},
				{0, 6},
				{2, 12},
				{11, 5},
				{6, 9},
				{10, 3},
			},
			patience: []int{
				0, 3, 2, 1, 5, 1, 5, 5, 3, 1, 2, 2, 2, 2, 1,
			},
			answer: 20,
		},
	}

	q := leetcode.Q2039{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.answer,
					q.NetworkBecomesIdle(tc.edges, tc.patience),
				)
			},
		)
	}
}
