package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	type testCase struct {
		nums []int
		ans  int
	}

	tcs := []testCase{
		{
			nums: []int{1, 1, 0, 1, 1, 1},
			ans:  3,
		},
		{
			nums: []int{1, 0, 1, 1, 0, 1},
			ans:  2,
		},
		{
			nums: []int{0, 0, 1, 0, 1, 1, 0, 1},
			ans:  2,
		},
		{
			nums: []int{0, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0},
			ans:  3,
		},
	}

	q := leetcode.Q0485{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.FindMaxConsecutiveOnes(tc.nums))
			},
		)
	}
}
