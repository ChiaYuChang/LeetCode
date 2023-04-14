package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	type testCase struct {
		words  []string
		answer []string
	}

	tcs := []testCase{
		{
			[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"},
			[]string{"catsdogcats", "dogcatsdog", "ratcatdogcat"},
		},
		{
			[]string{"cat", "dog", "catdog"},
			[]string{"catdog"},
		},
		{
			[]string{"t", "po", "nkq", "potn"},
			[]string{},
		},
		{
			[]string{"a", "b", "c", "cde", "d", "e", "abcde"},
			[]string{"cde", "abcde"},
		},
		{
			[]string{"mo", "pa", "molls", "molly", "pally"},
			[]string{},
		},
	}

	q := leetcode.Q0472{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-PrefixTree", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.ElementsMatch(t, tc.answer, q.FindAllConcatenatedWordsInADictsPrefixTree(tc.words))
			},
		)

		t.Run(
			fmt.Sprintf("Case %d-DP", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.ElementsMatch(t, tc.answer, q.FindAllConcatenatedWordsInADictsDP(tc.words))
			},
		)
	}
}
