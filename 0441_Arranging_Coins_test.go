package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestArrangeCoins(t *testing.T) {
	type testCase struct {
		n int
		a int
	}

	tcs := []testCase{
		{
			n: 1,
			a: 1,
		},
		{
			n: 5,
			a: 2,
		},
		{
			n: 8,
			a: 3,
		},
		{
			n: 1073741823,
			a: 46340,
		},
		{
			n: 2147483647,
			a: 65535,
		},
	}

	q := leetcode.Q0441{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%d", i+1, tc.n),
			func(t *testing.T) {
				require.Equal(t, tc.a, q.ArrangeCoins(tc.n))
				require.Equal(t, tc.a, q.ArrangeCoinsMathSqrt(tc.n))
			},
		)
	}
}
