package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestRemoveOuterParentheses(t *testing.T) {
	type testCase struct {
		s string
		a string
	}

	tcs := []testCase{
		{
			s: "(()())(())(()(()))",
			a: "()()()()(())",
		},
		{
			s: "(()())(())",
			a: "()()()",
		},
		{
			s: "()()",
			a: "",
		},
		{
			s: "",
			a: "",
		},
		{
			s: "((()())())(())",
			a: "(()())()()",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.a, leetcode.RemoveOuterParentheses(tc.s))
			},
		)
	}
}
