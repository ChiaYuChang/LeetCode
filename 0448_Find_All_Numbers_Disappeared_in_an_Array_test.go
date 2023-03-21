package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindDissappearedNumbers(t *testing.T) {
	type testCase struct {
		nums []int
		ans  []int
	}

	tcs := []testCase{
		{
			nums: []int{1},
			ans:  []int{},
		},
		{
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			ans:  []int{5, 6},
		},
		{
			nums: []int{1, 1},
			ans:  []int{2},
		},
		{
			nums: []int{
				30, 25, 26, 4, 21, 10, 28, 4,
				10, 7, 7, 19, 6, 26, 21, 10, 7,
				22, 9, 19, 7, 6, 26, 28, 1, 1,
				12, 27, 26, 17},
			ans: []int{
				2, 3, 5, 8, 11, 13, 14, 15,
				16, 18, 20, 23, 24, 29},
		},
	}

	q := leetcode.Q0448{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.ElementsMatch(
					t, tc.ans,
					q.FindDisappearedNumbers(tc.nums),
				)
				require.ElementsMatch(
					t, tc.ans,
					q.FindDisappearedNumbersSlow(tc.nums),
				)
			},
		)
	}
}
