package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinimumTime(t *testing.T) {
	type testCase struct {
		name       string
		time       []int
		totalTrips int
		ans        int64
	}

	tcs := []testCase{
		{
			name:       "Exact match",
			time:       []int{1, 2, 3},
			totalTrips: 5,
			ans:        3,
		},
		{
			name:       "length 1 time array",
			time:       []int{2},
			totalTrips: 1,
			ans:        2,
		},
		{
			name:       "0 total trips",
			time:       []int{1, 3, 5},
			totalTrips: 0,
			ans:        0,
		},
		{
			name:       "Not match",
			time:       []int{2, 4, 7, 11},
			totalTrips: 22,
			ans:        24,
		},
	}

	q := leetcode.Q2187{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name), func(t *testing.T) {
				require.Equal(t, tc.ans, q.MinimumTime(tc.time, tc.totalTrips))
			},
		)
	}
}
