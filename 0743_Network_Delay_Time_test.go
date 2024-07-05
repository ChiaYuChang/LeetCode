package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestNetworkDelayTime(t *testing.T) {
	type testCase struct {
		times [][]int
		n     int
		k     int
		ans   int
	}

	tcs := []testCase{
		{
			times: [][]int{
				{2, 1, 1},
				{2, 3, 1},
				{3, 4, 1},
			},
			n:   4,
			k:   2,
			ans: 2,
		},
		{
			times: [][]int{
				{1, 2, 1},
			},
			n:   2,
			k:   1,
			ans: 1,
		},
		{
			times: [][]int{
				{1, 2, 1},
			},
			n:   2,
			k:   2,
			ans: -1,
		},
	}

	q := leetcode.Q0743{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.NetworkDelayTime(tc.times, tc.n, tc.k))
			},
		)
	}
}
