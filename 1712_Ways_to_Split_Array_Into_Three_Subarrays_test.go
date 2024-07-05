package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestWaysToSplit(t *testing.T) {
	type testCase struct {
		num []int
		ans int
	}

	tcs := []testCase{
		{
			num: []int{1, 1, 1},
			ans: 1,
		},
		{
			num: []int{1, 2, 2, 2, 5, 0},
			ans: 3,
		},
		{
			num: []int{3, 2, 1},
			ans: 0,
		},
		{
			num: []int{1, 0, 0, 1, 0, 0, 1, 0},
			ans: 9,
		},
		{
			num: []int{0, 0, 0, 0, 1, 0},
			ans: 6,
		},
		{
			num: []int{0, 0, 0, 0, 0, 1, 2, 2, 2, 5, 0},
			ans: 28,
		},
		{
			num: []int{0, 0, 0, 0, 0, 0},
			ans: 10,
		},
		{
			num: []int{7, 2, 5, 5, 6, 2, 10, 9},
			ans: 6,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, leetcode.WaysToSplit(tc.num))
			},
		)
	}
}
