package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestIsValid(t *testing.T) {
	type testCase struct {
		str string
		ans bool
	}

	tcs := []testCase{
		{"()", true},
		{"{}()[]", true},
		{"{]", false},
		{"([)]", false},
	}

	q := leetcode.Q0020{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.str),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.IsValid(tc.str))
			},
		)
	}
}
