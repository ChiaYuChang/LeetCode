package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindDifference(t *testing.T) {
	type testCase struct {
		nums1, nums2 []int
		ans          [][]int
	}

	tcs := []testCase{
		{nums1: []int{1, 2, 3}, nums2: []int{2, 4, 6}, ans: [][]int{{1, 3}, {4, 6}}},
		{nums1: []int{1, 2, 3, 3}, nums2: []int{1, 1, 2, 2}, ans: [][]int{{3}, {}}},
	}

	q := leetcode.Q2215{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				mes := q.FindDifference(tc.nums1, tc.nums2)
				require.ElementsMatch(t, tc.ans[0], mes[0])
				require.ElementsMatch(t, tc.ans[1], mes[1])
			},
		)
	}
}
