package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestAddDigits(t *testing.T) {
	type testCase struct {
		num, ans int
	}

	tcs := []testCase{
		{38, 2},
		{18, 9},
		{0, 0},
		{10, 1},
		{9, 9},
	}

	q := leetcode.Q0258{}
	for i, tc := range tcs {
		t.Run(
			fmt.Sprintf("Case %d- Recursion %d", i, tc.num),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.AddDigits(tc.num))

			},
		)
		t.Run(
			fmt.Sprintf("Case %d- Math %d", i, tc.num),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.AddDigitsMath(tc.num))
			},
		)
	}
}
