package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSmallestNumber(t *testing.T) {
	type testCase struct {
		name string
		num  int64
		ans  int64
	}

	tcs := []testCase{
		{"Positive", 310, 103},
		{"Negative", -7605, -7650},
		{"Zero", 0, 0},
	}

	q := leetcode.Q2165{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.SmallestNumber(tc.num))
			},
		)
	}
}
