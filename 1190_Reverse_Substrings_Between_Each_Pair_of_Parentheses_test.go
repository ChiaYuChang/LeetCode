package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ1190(t *testing.T) {
	q := leetcode.Q1190{}

	tcs := []struct {
		str string
		ans string
	}{
		{"(abcd)", "dcba"},
		{"(u(love)i)", "iloveu"},
		{"(ed(et(oc))el)", "leetcode"},
		{"a(bc(de)fg(hi)j)k(lm)n", "ajhigfdecbkmln"},
		{"abcd((e))fg", "abcdefg"},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case: %s", tc.str),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.ReverseParentheses(tc.str))
			},
		)
	}
}
