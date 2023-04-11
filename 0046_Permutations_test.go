package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestPermute(t *testing.T) {
	type testCase struct {
		nums []int
		ans  [][]int
	}

	tcs := []testCase{
		{
			nums: []int{1, 2, 3},
			ans:  [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			nums: []int{0, 1},
			ans:  [][]int{{0, 1}, {1, 0}},
		},
		{
			nums: []int{0},
			ans:  [][]int{{0}},
		},
	}

	q := leetcode.Q0046{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.ElementsMatch(t, tc.ans, q.Permute(tc.nums))
			},
		)
	}
}
