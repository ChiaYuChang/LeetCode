package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestValidPerfectSquare(t *testing.T) {
	type testCase struct {
		num int
		ans bool
	}

	tcs := []testCase{
		{
			num: 76,
			ans: false,
		},
		{
			num: 3,
			ans: false,
		},
		{
			num: 1,
			ans: true,
		},
		{
			num: 178929,
			ans: true,
		},
		{
			num: 178930,
			ans: false,
		},
		{
			num: 25,
			ans: true,
		},
	}

	q := leetcode.Q0367{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d %d", i+1, tc.num),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.IsPerfectSquare(tc.num))
			})
	}
}
