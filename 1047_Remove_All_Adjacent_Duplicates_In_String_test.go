package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		str string
		ans string
	}

	tcs := []testCase{
		{"abbaca", "ca"},
		{"azxxzy", "ay"},
		{"abbbc", "abc"},
	}

	q := leetcode.Q1047{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.str),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.RemoveDuplicates(tc.str))
			},
		)
	}
}
