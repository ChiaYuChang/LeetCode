package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestDecimalToHex(t *testing.T) {
	type testCase struct {
		num int
		ans string
	}

	tcs := []testCase{
		{
			num: 26,
			ans: "1a",
		},
		{
			num: -1,
			ans: "ffffffff",
		},
		{
			num: -26,
			ans: "ffffffe6",
		},
		{
			num: -26241,
			ans: "ffff997f",
		},
		{
			num: 224243,
			ans: "36bf3",
		},
		{
			num: 0,
			ans: "0",
		},
	}

	d2h := leetcode.Q0405DecimalToHex{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d %d", i+1, tc.num),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans,
					d2h.ToHex(tc.num))
			},
		)
	}
}
