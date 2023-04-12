package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLongestPalindromeSubseq(t *testing.T) {
	type testCase struct {
		str string
		ans int
	}

	tcs := []testCase{
		{"bbbsb", 4},
		{"cbbd", 2},
		{"abaabaa", 6},
	}

	q := leetcode.Q0516{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.LongestPalindromeSubseq(tc.str))
			},
		)
	}
}
