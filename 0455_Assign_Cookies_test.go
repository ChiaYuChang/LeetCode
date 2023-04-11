package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindContentChildren(t *testing.T) {
	type testCase struct {
		g, s []int
		ans  int
	}

	tcs := []testCase{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}

	q := leetcode.Q0455{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.FindContentChildren(tc.g, tc.s))
			},
		)
	}
}
