package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinCostClimbingStairs(t *testing.T) {
	type testCase struct {
		name   string
		cost   []int
		answer int
	}

	tcs := []testCase{
		{
			name:   "Odd I",
			cost:   []int{10, 15, 20},
			answer: 15,
		},
		{
			name:   "Odd II",
			cost:   []int{10, 15, 20, 17, 13},
			answer: 32,
		},
		{
			name:   "Even I",
			cost:   []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			answer: 6,
		},
		{
			name:   "Even II",
			cost:   []int{10, 15, 20, 17, 13, 20},
			answer: 43,
		},
	}

	q := leetcode.Q0746{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				require.Equal(t, tc.answer, q.MinCostClimbingStairs(tc.cost))
			},
		)
	}
}
