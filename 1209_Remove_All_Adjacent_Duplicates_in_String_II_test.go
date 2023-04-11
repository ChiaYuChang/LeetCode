package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestRemoveDuplicatesII(t *testing.T) {
	type testCase struct {
		s string
		k int
		a string
	}

	tcs := []testCase{
		{"abcd", 2, "abcd"},
		{"deeedbbcccbdaa", 3, "aa"},
		{"pbbcggttciiippooaais", 2, "ps"},
	}

	q := leetcode.Q1209{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.s),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.a, q.RemoveDuplicates(tc.s, tc.k))
			},
		)
	}
}
