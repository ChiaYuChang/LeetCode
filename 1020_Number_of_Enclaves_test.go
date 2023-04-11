package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestNumEnclaves(t *testing.T) {
	type testCase struct {
		grid [][]int
		ans  int
	}

	tcs := []testCase{
		{
			grid: [][]int{
				{0, 0, 0, 0},
				{1, 0, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			ans: 3,
		},
		{
			grid: [][]int{
				{0, 1, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			ans: 0,
		},
		{
			grid: [][]int{{1}},
			ans:  0,
		},
		{
			grid: [][]int{{0}, {1}, {1}, {0}, {0}},
			ans:  0,
		},
		{
			grid: [][]int{
				{0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1},
				{1, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1},
				{1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 0},
				{1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1},
				{1, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0},
				{1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0},
			},
			ans: 27,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, leetcode.NumEnclaves(tc.grid))
			},
		)
	}
}
