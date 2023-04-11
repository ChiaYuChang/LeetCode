package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinPath(t *testing.T) {
	type testCase struct {
		grid [][]int
		ans  int
	}

	tcs := []testCase{
		{
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			ans: 7,
		},
		{
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			ans: 12,
		},
		{
			grid: [][]int{{1}},
			ans:  1,
		},
	}

	q := leetcode.Q0064{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.MinPathSum(tc.grid))
			},
		)
	}
}
