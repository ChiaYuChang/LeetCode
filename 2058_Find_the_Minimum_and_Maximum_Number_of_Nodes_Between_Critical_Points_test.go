package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ2058(t *testing.T) {
	tcs := []struct {
		name string
		nums []int
		ans  []int
	}{
		{
			name: "General Case 1",
			nums: []int{5, 3, 1, 2, 5, 1, 2},
			ans:  []int{1, 3},
		},
		{
			name: "General Case 2",
			nums: []int{1, 3, 2, 2, 3, 2, 2, 2, 7},
			ans:  []int{3, 3},
		},
		{
			name: "List with less than 3 elements",
			nums: []int{3, 1},
			ans:  []int{-1, -1},
		},
		{
			name: "No critical points",
			nums: []int{2, 3, 3, 2},
			ans:  []int{-1, -1},
		},
	}

	q := leetcode.Q2058{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				t.Parallel()
				result := q.NodesBetweenCriticalPoints(q.NewListNodeFromList(tc.nums))
				require.Equal(t, tc.ans, result)
			},
		)
	}
}
