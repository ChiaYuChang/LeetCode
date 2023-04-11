package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestPalindromePartition(t *testing.T) {
	type testCase struct {
		str string
		ans [][]string
	}

	tcs := []testCase{
		{
			"aab",
			[][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			"a",
			[][]string{
				{"a"},
			},
		},
		{
			"abcdbbfcba",
			[][]string{
				{"a", "b", "c", "d", "bb", "f", "c", "b", "a"},
				{"a", "b", "c", "d", "b", "b", "f", "c", "b", "a"},
			},
		},
		{
			"abba",
			[][]string{
				{"abba"},
				{"a", "bb", "a"},
				{"a", "b", "b", "a"},
			},
		},
		{
			"abcba",
			[][]string{
				{"abcba"},
				{"a", "bcb", "a"},
				{"a", "b", "c", "b", "a"},
			},
		},
		{
			"abadddabae",
			[][]string{
				{"abadddaba", "e"},
				{"aba", "ddd", "aba", "e"},
				{"aba", "ddd", "a", "b", "a", "e"},
				{"aba", "dd", "d", "aba", "e"},
				{"aba", "dd", "d", "a", "b", "a", "e"},
				{"aba", "d", "dd", "aba", "e"},
				{"aba", "d", "dd", "a", "b", "a", "e"},
				{"aba", "d", "d", "d", "aba", "e"},
				{"aba", "d", "d", "d", "a", "b", "a", "e"},
				{"a", "badddab", "a", "e"},
				{"a", "b", "addda", "b", "a", "e"},
				{"a", "b", "a", "ddd", "aba", "e"},
				{"a", "b", "a", "ddd", "a", "b", "a", "e"},
				{"a", "b", "a", "dd", "d", "aba", "e"},
				{"a", "b", "a", "dd", "d", "a", "b", "a", "e"},
				{"a", "b", "a", "d", "dd", "aba", "e"},
				{"a", "b", "a", "d", "dd", "a", "b", "a", "e"},
				{"a", "b", "a", "d", "d", "d", "aba", "e"},
				{"a", "b", "a", "d", "d", "d", "a", "b", "a", "e"},
			},
		},
	}

	q := leetcode.Q0131{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Cast %d-%s", i+1, tc.str),
			func(t *testing.T) {
				t.Parallel()
				require.ElementsMatch(t, tc.ans, q.Partition(tc.str))
			},
		)
	}
}
