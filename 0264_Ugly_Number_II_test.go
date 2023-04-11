package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestNthUglyNumber(t *testing.T) {
	type testCase struct {
		num int
		ans int
	}

	tcs := []testCase{
		{10, 12},
		{1, 1},
	}

	q := leetcode.Q0264{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%d", i+1, tc.num),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.SlowNthUglyNumberII(tc.num))
			},
		)
	}
}
