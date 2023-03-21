package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestNumberOfDigitOne(t *testing.T) {
	type testCase struct {
		name string
		num  int
		ans  int
	}

	tcs := []testCase{
		{
			name: "General",
			num:  13,
			ans:  6,
		},
		{
			name: "Lower bound",
			num:  0,
			ans:  0,
		},
		{
			name: "Upper bound",
			num:  1000000000,
			ans:  900000001,
		},
		{
			name: "Random large number",
			num:  1412731,
			ans:  1321318,
		},
	}

	q := leetcode.Q0233{}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.NumberOfDigitOne(tc.num))
			},
		)
	}
}
