package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinEatingSpeed(t *testing.T) {
	type testCase struct {
		piles    []int
		hours    int
		minSpeed int
	}

	tcs := []testCase{
		{
			piles:    []int{3, 6, 7, 11},
			hours:    8,
			minSpeed: 4,
		},
		{
			piles:    []int{30, 11, 23, 4, 20},
			hours:    5,
			minSpeed: 30,
		},
		{
			piles:    []int{30, 11, 23, 4, 20},
			hours:    6,
			minSpeed: 23,
		},
	}

	q := leetcode.Q0875{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.minSpeed, q.MinEatingSpeed(tc.piles, tc.hours))
			},
		)
	}
}
