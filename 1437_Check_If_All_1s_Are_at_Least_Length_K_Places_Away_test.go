package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestKLengthApart(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		ans  bool
	}

	tcs := []testCase{
		{
			nums: []int{1, 0, 0, 0, 1, 0, 0, 1},
			k:    2,
			ans:  true,
		},
		{
			nums: []int{1, 0, 0, 1, 0, 1},
			k:    2,
			ans:  false,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, leetcode.KLengthApart(tc.nums, tc.k))
			},
		)
	}
}
