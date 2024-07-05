package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindMiddleIndex(t *testing.T) {
	type testCase struct {
		nums []int
		ans  int
	}

	tcs := []testCase{
		{[]int{2, 3, -1, 8, 4}, 3},
		{[]int{1, -1, 4}, 2},
		{[]int{2, 5}, -1},
	}

	q := leetcode.Q1991{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.FindMiddleIndex(tc.nums))
			},
		)
	}
}
