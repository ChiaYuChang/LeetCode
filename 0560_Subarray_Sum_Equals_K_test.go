package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSubarraySum(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		ans  int
	}

	tcs := []testCase{
		// {[]int{1, 1, 1}, 2, 2},
		// {[]int{1, 2, 3}, 3, 2},
		{[]int{-1, 1, 0, 1, -1}, 0, 6},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, leetcode.SubarraySum(tc.nums, tc.k))
			},
		)
	}
}
