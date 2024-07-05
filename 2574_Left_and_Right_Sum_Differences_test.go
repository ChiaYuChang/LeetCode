package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLeftRigthDifference(t *testing.T) {
	type testCase struct {
		name string
		nums []int
		ans  []int
	}

	tcs := []testCase{
		{
			name: "General case",
			nums: []int{10, 4, 8, 3},
			ans:  []int{15, 1, 11, 22},
		},
		{
			name: "Single element",
			nums: []int{1},
			ans:  []int{0},
		},
	}

	q := leetcode.Q2574{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.LeftRigthDifference(tc.nums))
			},
		)
	}
}
