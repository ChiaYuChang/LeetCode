package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestTopKFreqElements(t *testing.T) {
	type testCase struct {
		name string
		nums []int
		k    int
		ans  []int
	}

	tcs := []testCase{
		{
			name: "General",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			ans:  []int{1, 2},
		},
		{
			name: "Lenght 1 nums",
			nums: []int{1},
			k:    1,
			ans:  []int{1},
		},
	}

	q := leetcode.Q0347{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				require.ElementsMatch(t, tc.ans, q.TopKFrequent(tc.nums, tc.k))
			},
		)
	}
}
