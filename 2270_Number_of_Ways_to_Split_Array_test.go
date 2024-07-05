package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestWaysToSplitArray(t *testing.T) {
	type testCase struct {
		nums []int
		ans  int
	}

	tcs := []testCase{
		{[]int{10, 4, -8, 7}, 2},
		{[]int{2, 3, 1, 0}, 2},
		{[]int{0, 0}, 1},
	}

	q := leetcode.Q2270{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.WaysToSplitArray(tc.nums))
			},
		)
	}
}
