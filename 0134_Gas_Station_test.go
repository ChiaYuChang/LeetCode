package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestCanCompleteCircuit(t *testing.T) {
	type testCase struct {
		gas  []int
		cost []int
		ans  int
	}

	tcs := []testCase{
		{
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			ans:  3,
		},
		{
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			ans:  -1,
		},
		{
			gas:  []int{4, 5, 2, 6, 5, 3},
			cost: []int{3, 2, 7, 3, 2, 9},
			ans:  -1,
		},
		{
			gas:  []int{5, 8, 2, 8},
			cost: []int{6, 5, 6, 6},
			ans:  3,
		},
		{
			gas:  []int{2, 8, 9, 1, 2, 7, 4, 4, 1, 6, 3},
			cost: []int{3, 6, 7, 6, 3, 5, 2, 9, 2, 1, 3},
			ans:  9,
		},
		{
			gas:  []int{3},
			cost: []int{3},
			ans:  0,
		},
		{
			gas:  []int{2},
			cost: []int{3},
			ans:  -1,
		},
	}

	q := leetcode.Q0134{}
	for i := range tcs {
		tc := tcs[i]
		require.Equal(t, tc.ans,
			q.CanCompleteCircuit(tc.gas, tc.cost))
		require.Equal(t, tc.ans,
			q.SlowCanCompleteCircuit(tc.gas, tc.cost))
	}
}
