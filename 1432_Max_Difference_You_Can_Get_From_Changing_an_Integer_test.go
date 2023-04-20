package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMaxDiff(t *testing.T) {
	type testCase struct {
		name string
		num  int
		ans  int
	}

	tcs := []testCase{
		{"General Case", 555, 888},
		{"num is Less than 10", 9, 8},
		{"1XX", 199, 899},
		{"99X", 990, 889},
		{"10X", 1101057, 8808050},
	}

	q := leetcode.Q1432{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.MaxDiff(tc.num))
			},
		)
	}
}
