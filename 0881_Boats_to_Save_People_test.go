package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestNumRescueBoats(t *testing.T) {
	type testCase struct {
		people []int
		limit  int
		nBoat  int
	}

	tcs := []testCase{
		{[]int{1, 2}, 3, 1},
		{[]int{3, 2, 2, 1}, 3, 3},
		{[]int{3, 5, 3, 4}, 5, 4},
		{[]int{3, 2, 3, 2, 2}, 6, 3},
	}

	q := leetcode.Q0881{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.nBoat, q.NumRescueBoats(tc.people, tc.limit))
			},
		)
	}
}
