package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestPermuteII(t *testing.T) {

	type testCase struct {
		nums []int
		n    int
		ans  [][]int
	}

	tcs := []testCase{
		{
			nums: []int{0, 1, 2, 3},
			n:    2,
			ans: [][]int{
				{0, 1}, {0, 2}, {0, 3},
				{1, 0}, {1, 2}, {1, 3},
				{2, 0}, {2, 1}, {2, 3},
				{3, 0}, {3, 1}, {3, 2},
			},
		},
		{
			nums: []int{0, 1, 2, 3, 4, 5, 6},
			n:    1,
			ans: [][]int{
				{0}, {1}, {2}, {3},
				{4}, {5}, {6},
			},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d n = %d", i+1, tc.n),
			func(t *testing.T) {
				require.ElementsMatch(t, tc.ans, leetcode.Permute(tc.nums, tc.n))
			},
		)
	}
}
