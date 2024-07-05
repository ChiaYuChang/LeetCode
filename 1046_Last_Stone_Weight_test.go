package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLastStoneWeight(t *testing.T) {
	type testCase struct {
		stones []int
		answer int
	}

	tcs := []testCase{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1}, 1},
	}

	q := leetcode.Q1046{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.answer, q.LastStoneWeight(tc.stones))
			},
		)
	}
}
