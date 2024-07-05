package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestIntersectionOfMutipleSets(t *testing.T) {
	type testCase struct {
		nums [][]int
		ans  []int
	}

	tcs := []testCase{
		{
			nums: [][]int{
				{3, 1, 2, 4, 5},
				{1, 2, 3, 4},
				{3, 4, 5, 6},
			},
			ans: []int{
				3, 4,
			},
		},
		{
			nums: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			ans: []int{},
		},
		{
			nums: [][]int{
				{4, 2, 3, 1},
			},
			ans: []int{
				1, 2, 3, 4,
			},
		},
	}

	q := leetcode.Q2248{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.Intersection(tc.nums))
			},
		)
	}
}
