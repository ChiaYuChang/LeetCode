package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestKSmallestPairs(t *testing.T) {
	type testCase struct {
		nums1, nums2 []int
		k            int
		ans          [][]int
	}

	tcs := []testCase{
		{
			nums1: []int{1, 7, 11},
			nums2: []int{2, 4, 6},
			k:     3,
			ans:   [][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			nums1: []int{1, 1, 2},
			nums2: []int{1, 2, 3},
			k:     2,
			ans:   [][]int{{1, 1}, {1, 1}},
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3},
			k:     3,
			ans:   [][]int{{1, 3}, {2, 3}},
		},
		{
			nums1: []int{1, 1, 2},
			nums2: []int{1, 2, 3},
			k:     10,
			ans:   [][]int{{1, 1}, {1, 1}, {1, 2}, {1, 2}, {2, 1}, {2, 2}, {1, 3}, {1, 3}, {2, 3}},
		},
	}

	q := leetcode.Q0373{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.ElementsMatch(t, tc.ans, q.KSmallestPairs(tc.nums1, tc.nums2, tc.k))
			},
		)
	}
}
