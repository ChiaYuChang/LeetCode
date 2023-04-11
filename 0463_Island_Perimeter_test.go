package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestIslandPerimeter(t *testing.T) {
	type testCase struct {
		grid [][]int
		ans  int
	}

	tcs := []testCase{
		{
			grid: [][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			ans: 16,
		},
		{
			grid: [][]int{
				{1},
			},
			ans: 4,
		},
		{
			grid: [][]int{
				{0, 1},
			},
			ans: 4,
		},
	}

	q := leetcode.Q0463{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.IslandPerimeter(tc.grid))
			},
		)
	}
}
