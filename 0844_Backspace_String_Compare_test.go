package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestBackspaceCompare(t *testing.T) {
	type testCase struct {
		s1, s2 string
		ans    bool
	}

	tcs := []testCase{
		{"ab#c", "ac#c", true},
		{"ab##", "c#d#", true},
		{"a#c", "b", false},
	}

	q := leetcode.Q0844{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.BackspaceCompare(tc.s1, tc.s2))
			},
		)
	}
}
