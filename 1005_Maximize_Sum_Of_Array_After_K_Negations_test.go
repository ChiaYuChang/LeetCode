package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLargestSumAfterKNegations(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		ans  int
	}
	tcs := []testCase{
		{[]int{4, 2, 3}, 1, 5},
		{[]int{3, -1, 0, 2}, 3, 6},
		{[]int{2, -3, -1, 5, -4}, 2, 13},
		{[]int{-4, -2, -3}, 4, 5},
		{[]int{0, 0, 0, 0, 0, 0, 0}, 100, 0},
	}
	q := leetcode.Q1005{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.LargestSumAfterKNegations(tc.nums, tc.k))
			},
		)
	}
}
