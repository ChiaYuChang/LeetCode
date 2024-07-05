package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ1509(t *testing.T) {
	q := leetcode.Q1509{}

	tcs := []struct {
		nums []int
		ans  int
	}{
		{[]int{5, 3, 2, 4}, 0},
		{[]int{1, 5, 0, 10, 14}, 1},
		{[]int{3, 100, 20}, 0},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case_%d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.MinDifference(tc.nums))
			})
	}
}
