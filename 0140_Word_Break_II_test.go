package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestWordBreak(t *testing.T) {
	type testCase struct {
		s        string
		wordDict []string
		answer   []string
	}

	tcs := []testCase{
		{
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			answer:   []string{"cats and dog", "cat sand dog"},
		},
		{
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			answer:   []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			answer:   []string{},
		},
		{
			s:        "ab",
			wordDict: []string{"a", "b"},
			answer:   []string{"a b"},
		},
	}

	q := leetcode.Q0140{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.s),
			func(t *testing.T) {
				t.Parallel()
				require.ElementsMatch(t, tc.answer, q.WordBreak(tc.s, tc.wordDict))
			},
		)
	}
}
