package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestTwoSum(t *testing.T) {

	type testCase struct {
		nums   []int
		target int
		ans    []int
	}

	tcs := []testCase{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			ans:    []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			ans:    []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			ans:    []int{0, 1},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i),
			func(t *testing.T) {
				t.Parallel()
				rst := leetcode.TwoSum(tc.nums, tc.target)
				require.Len(t, rst, 2)
				if rst[0] > rst[1] {
					rst[0], rst[1] = rst[1], rst[0]
				}
				require.Equal(t, tc.ans[0], rst[0])
				require.Equal(t, tc.ans[1], rst[1])
			},
		)
	}
}
