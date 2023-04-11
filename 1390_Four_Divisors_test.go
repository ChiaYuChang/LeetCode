package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSumFourDivisors(t *testing.T) {
	type testCase struct {
		nums []int
		ans  int
	}

	tcs := []testCase{
		{
			nums: []int{21, 4, 7},
			ans:  32,
		},
		{
			nums: []int{21, 21},
			ans:  64,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			ans:  0,
		},
	}

	q := leetcode.Q1390{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.SumFourDivisors(tc.nums))
			},
		)
	}
}
