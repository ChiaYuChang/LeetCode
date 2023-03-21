package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSumRange(t *testing.T) {
	type testCase struct {
		steps  []string
		params [][]int
		answer []int
	}

	tcs := []testCase{
		{
			steps: []string{
				"NumArray", "sumRange", "sumRange", "sumRange",
			},
			params: [][]int{
				{-2, 0, 3, -5, 2, -1},
				{0, 2},
				{2, 5},
				{0, 5},
			},
			answer: []int{
				0, 1, -1, -3,
			},
		},
		{
			steps: []string{
				"NumArray", "sumRange", "sumRange", "sumRange", "sumRange", "sumRange",
			},
			params: [][]int{
				{31, 23, 2, -2, 10, 3, -25, 23, -1, -23, 80},
				{0, 2},
				{2, 5},
				{0, 5},
				{3, 9},
				{8, 10},
			},
			answer: []int{
				0, 56, 13, 67, -15, 56,
			},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				q := leetcode.Q0303{}

				var na leetcode.Q0303NumArray
				for i, s := range tc.steps {
					switch s {
					case "NumArray":
						na = q.NumArrayConstructor(tc.params[i])
					case "sumRange":
						require.Equal(t,
							tc.answer[i],
							na.SumRange(tc.params[i][0], tc.params[i][1]),
						)
					}
				}
			})
	}
}
