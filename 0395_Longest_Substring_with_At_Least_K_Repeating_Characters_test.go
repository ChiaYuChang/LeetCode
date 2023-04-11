package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLongestSubstring(t *testing.T) {
	type testCase struct {
		s string
		k int
		a int
	}

	tcs := []testCase{
		{"aaabb", 3, 3},
		{"ababbc", 2, 5},
		{"ababacb", 3, 0},
		{"aaabbb", 3, 6},
		{"aaaaa", 6, 0},
	}

	q := leetcode.Q0395{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-Recursive", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.a, q.LongestSubstringRecursive(tc.s, tc.k))
			},
		)
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-Sliding Windows", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.a, q.LongestSubstringSlidingWindows(tc.s, tc.k))
			},
		)
	}
}
