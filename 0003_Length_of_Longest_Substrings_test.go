package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLenghOfLongestSubstring(t *testing.T) {
	type testCase struct {
		n string
		s string
		l int
	}

	tcs := []testCase{
		{
			n: "Case_1",
			s: "abcabcbb",
			l: 3,
		},
		{
			n: "Case_2",
			s: "bbbbb",
			l: 1,
		},
		{
			n: "Case_3",
			s: "pwwkew",
			l: 3,
		},
		{
			n: "Non-repeated string",
			s: "abcd",
			l: 4,
		},
		{
			n: "Length 1 string",
			s: "a",
			l: 1,
		},
		{
			n: "Empty string",
			s: "",
			l: 0,
		},
		{
			n: "Repeat at the end",
			s: "abb",
			l: 2,
		},
		{
			n: "Repeat at the frist two",
			s: "aab",
			l: 2,
		},
		{
			n: "Odd Palindrome",
			s: "abcba",
			l: 3,
		},
		{
			n: "Even Palindrome",
			s: "abba",
			l: 2,
		},
	}

	q := leetcode.Q0003{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.n,
			func(t *testing.T) {
				t.Parallel()
				ans := q.LengthOfLongestSubstring(tc.s)
				require.Equal(t, tc.l, ans)
			},
		)
	}
}
