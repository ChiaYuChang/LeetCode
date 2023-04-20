package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMaxSubsequence(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		ans  []int
	}

	tcs := []testCase{
		{
			nums: []int{2, 1, 3, 3},
			k:    2,
			ans:  []int{3, 3},
		},
		{
			nums: []int{-1, -2, 3, 4},
			k:    3,
			ans:  []int{-1, 3, 4},
		},
		{
			nums: []int{3, 4, 3, 3},
			k:    2,
			ans:  []int{3, 4},
		},
		{
			nums: []int{30, -10},
			k:    2,
			ans:  []int{30, -10},
		},
	}

	q := leetcode.Q2099{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.MaxSubsequence(tc.nums, tc.k))
			},
		)
	}
}
