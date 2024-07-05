package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFourSumCount(t *testing.T) {
	type testCase struct {
		num1, num2, num3, num4 []int
		ans                    int
	}

	tcs := []testCase{
		{
			num1: []int{1, 2},
			num2: []int{-2, -1},
			num3: []int{-1, 2},
			num4: []int{0, 2},
			ans:  2,
		},
		{
			num1: []int{0},
			num2: []int{0},
			num3: []int{0},
			num4: []int{0},
			ans:  1,
		},
		{
			num1: []int{-1, -1},
			num2: []int{-1, 1},
			num3: []int{-1, 1},
			num4: []int{1, -1},
			ans:  6,
		},
		{
			num1: []int{1, 1, -1, -1},
			num2: []int{-1, -1, 1, 1},
			num3: []int{1, -1, 0, -1},
			num4: []int{1, 1, -1, 1},
			ans:  76,
		},
	}

	q := leetcode.Q0454{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.FourSumCount(
					tc.num1, tc.num2, tc.num3, tc.num4,
				))
			},
		)
	}
}
