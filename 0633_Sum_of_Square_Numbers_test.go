package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSumOfSqureNumbers(t *testing.T) {
	type testCase struct {
		num int
		ans bool
	}

	tcs := []testCase{
		{
			num: 0,
			ans: true,
		},
		{
			num: 1,
			ans: true,
		},
		{
			num: 2,
			ans: true,
		},
		{
			num: 689,
			ans: true,
		},
		{
			num: 690,
			ans: false,
		},
	}

	q := leetcode.Q0633{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %2d %d", i+1, tc.num),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.JudgeSquareSum(tc.num))
			},
		)
	}
}
