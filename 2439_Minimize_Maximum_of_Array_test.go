package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinimizeArrayValue(t *testing.T) {
	type testCase struct {
		nums   []int
		minmax int
	}

	tcs := []testCase{
		{
			nums:   []int{3, 7, 1, 6},
			minmax: 5,
		},
		{
			nums:   []int{10, 1},
			minmax: 10,
		},
	}

	q := leetcode.Q2439{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.minmax, q.MinimizeArrayValue(tc.nums))
			},
		)
	}
}
