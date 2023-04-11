package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMaxRotateFunction(t *testing.T) {
	type testCase struct {
		name string
		nums []int
		ans  int
	}

	tcs := []testCase{
		{
			name: "General case",
			nums: []int{4, 3, 2, 6},
			ans:  26,
		},
		{
			name: "Single element",
			nums: []int{100},
			ans:  0,
		},
		{
			name: "With negative valuse",
			nums: []int{4, -3, 2, 6, -10, 12},
			ans:  52,
		},
	}

	q := leetcode.Q0396{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.MaxRotateFunction(tc.nums))
			},
		)
	}
}
