package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinCut(t *testing.T) {
	type testCase struct {
		s string
		a int
	}

	tcs := []testCase{
		{"aab", 1},
		{"a", 0},
		{"ab", 1},
		{"ababababababababababababcbabababababababababababa", 0},
		{"ccaacabacb", 3},
	}

	q := leetcode.Q0132{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.a, q.MinCut(tc.s))
			},
		)
	}
}
