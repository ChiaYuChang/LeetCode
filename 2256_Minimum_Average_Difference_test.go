package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinimumAverageDifference(t *testing.T) {
	type testCase struct {
		nums []int
		ans  int
	}

	tcs := []testCase{
		{
			nums: []int{0},
			ans:  0,
		},
		{
			nums: []int{4, 2, 0},
			ans:  2,
		},
		{
			nums: []int{0, 0, 0, 0, 0},
			ans:  0,
		},
		{
			nums: []int{0, 1, 0, 1, 0, 1},
			ans:  0,
		},
	}

	q := leetcode.Q2256{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.MinimumAverageDifference(tc.nums))
			},
		)
	}
}
