package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSortByBits(t *testing.T) {
	type testCase struct {
		arr []int
		ans []int
	}
	tcs := []testCase{
		{
			arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			ans: []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			arr: []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			ans: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, leetcode.SortByBits(tc.arr))
			},
		)
	}
}
