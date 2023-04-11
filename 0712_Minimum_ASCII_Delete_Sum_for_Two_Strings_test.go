package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinimumDeleteSum(t *testing.T) {
	type testCase struct {
		s1, s2 string
		answer int
	}

	tcs := []testCase{
		{"sea", "eat", 231},
		{"delete", "leet", 403},
		{"a", "", 97},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.answer, leetcode.MinimumDeleteSum(tc.s1, tc.s2))
			},
		)
	}
}
