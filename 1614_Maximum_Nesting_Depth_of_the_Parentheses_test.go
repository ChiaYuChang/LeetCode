package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMaxDepth(t *testing.T) {
	type testCase struct {
		s string
		a int
	}

	tcs := []testCase{
		{"(1+(2*3)+((8)/4))+1", 3},
		{"(1)+((2))+(((3)))", 3},
		{"1", 0},
	}

	q := leetcode.Q1614{}
	for i, tc := range tcs {
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.a, q.MaxDepth(tc.s))
			},
		)
	}
}
