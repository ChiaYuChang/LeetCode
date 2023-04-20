package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMergeAlternately(t *testing.T) {
	type testCase struct {
		word1, word2 string
		answer       string
	}

	tcs := []testCase{
		{word1: "abc", word2: "pqr", answer: "apbqcr"},
		{word1: "ab", word2: "pqrs", answer: "apbqrs"},
		{word1: "abcd", word2: "pq", answer: "apbqcd"},
	}

	q := leetcode.Q1768{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.answer, q.MergeAlternately(tc.word1, tc.word2))
			},
		)
	}
}
