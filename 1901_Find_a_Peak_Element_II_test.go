package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindAPeakElementII(t *testing.T) {
	type testCase struct {
		name string
		ary  []int
		ans  int
	}

	tcs := []testCase{
		{
			name: "Case 1",
			ary:  []int{0, 1, 0},
			ans:  1,
		},
		{
			name: "Case 2",
			ary:  []int{0, 2, 1, 0},
			ans:  1,
		},
		{
			name: "Case 3",
			ary:  []int{0, 10, 5, 2},
			ans:  1,
		},
		{
			name: "Case 4",
			ary:  []int{8, 18, 24, 31, 37, 42, 43, 56, 65, 73, 93, 98, 100, 98, 76, 72, 69, 24, 23},
			ans:  12,
		},
	}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				t.Parallel()
				require.Equal(
					t, tc.ans,
					leetcode.PeakIndexInMountainArray(tc.ary))
			},
		)
	}
}
