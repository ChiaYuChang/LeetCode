package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestJumpGameII(t *testing.T) {
	type testCase struct {
		nums []int
		ans  int
	}

	tcs := []testCase{
		{
			nums: []int{2, 3, 1, 1, 4},
			ans:  2,
		},
		{
			nums: []int{2, 3, 0, 1, 4},
			ans:  2,
		},
	}

	q := leetcode.Q0045{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d by DB O(n^2)", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.JumpSlow(tc.nums))
			},
		)
		t.Run(
			fmt.Sprintf("Case %d by O(n)", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.JumpFast(tc.nums))
			},
		)
	}
}
