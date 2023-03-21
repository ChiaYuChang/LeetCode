package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	type testCase struct {
		str string
		ans bool
	}

	tcs := []testCase{
		{"abab", true},
		{"aba", false},
		{"abcabcabcabc", true},
		{"bb", true},
	}

	q := leetcode.Q0459{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(
					t, tc.ans,
					q.RepeatedSubstringPattern(tc.str),
				)
			},
		)
	}
}
