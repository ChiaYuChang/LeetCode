package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func TestLongestPalindrome(t *testing.T) {
	type testCase struct {
		name string
		s    string
		ans  string
	}

	tcs := []testCase{
		{
			name: "Case_1",
			s:    "abcdefgbabad",
			ans:  "aba",
		},
		{
			name: "Case_2",
			s:    "cbbd",
			ans:  "bb",
		},
		{
			name: "Case_3",
			s:    "a",
			ans:  "a",
		},
		{
			name: "Case_5",
			s:    "accbdkcca",
			ans:  "cc",
		},
		{
			name: "Case_6",
			s:    "aacabdkacaa",
			ans:  "aca",
		},
		{
			name: "Case_7",
			s:    "oeabcdbbfcbak",
			ans:  "bb",
		},
		{
			name: "Case_8",
			s:    "abcdbbfcba",
			ans:  "bb",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				p := leetcode.LongestPalindrome(tc.s)
				require.Equal(t, tc.ans, p)
				require.True(t, isPalindrome(p))
			},
		)
	}
}
