package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestRepeatedStringMatch(t *testing.T) {
	type testCase struct {
		name string
		a, b string
		ans  int
	}

	tcs := []testCase{
		{
			name: "OK",
			a:    "a",
			b:    "aa",
			ans:  2,
		},
		{
			name: "Incompleted start",
			a:    "abcd",
			b:    "cdabcdab",
			ans:  3,
		},
		{
			name: "len(a) > len(b)",
			a:    "aa",
			b:    "a",
			ans:  1,
		},
		{
			name: "len(a) > len(b), incompleted start",
			a:    "aaaaab",
			b:    "aabaa",
			ans:  2,
		},
		{
			name: "Not OK",
			a:    "aaab",
			b:    "abaaabaaabaacbaaab",
			ans:  -1,
		},
	}

	q := leetcode.Q0686{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.RepeatedStringMatch(tc.a, tc.b))
			},
		)
	}
}
