package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLongestCommonSubsequence(t *testing.T) {
	type testCase struct {
		text1, text2 string
		cssLen       int
	}

	tcs := []testCase{
		{"abcde", "acd", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}

	q := leetcode.Q1143{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.cssLen, q.LongestCommonSubsequence(tc.text1, tc.text2))
			},
		)
	}
}
