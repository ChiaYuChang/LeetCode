package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSlidingWindowMax(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		ans  []int
	}

	tcs := []testCase{
		{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			ans:  []int{3, 3, 5, 5, 6, 7},
		},
		{
			nums: []int{1},
			k:    1,
			ans:  []int{1},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case_%d", i),
			func(t *testing.T) {
				m := leetcode.MaxSlidingWindow(tc.nums, tc.k)
				require.Equal(t, len(tc.ans), len(m))
				for i, a := range tc.ans {
					require.Equal(t, a, tc.ans[i])
				}
			},
		)
	}
}
