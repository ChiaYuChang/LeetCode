package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMakeConnected(t *testing.T) {
	type testCase struct {
		n           int
		connections [][]int
		ans         int
	}

	tcs := []testCase{
		{
			n: 4,
			connections: [][]int{
				{0, 1},
				{0, 2},
				{1, 2},
			},
			ans: 1,
		},
		{
			n: 6,
			connections: [][]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{1, 2},
				{1, 3},
			},
			ans: 2,
		},
		{
			n: 6,
			connections: [][]int{
				{0, 1},
				{0, 2},
				{0, 3},
				{1, 2},
			},
			ans: -1,
		},
	}

	q := leetcode.Q1319{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case: %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.MakeConnected(tc.n, tc.connections))
			},
		)
	}
}
