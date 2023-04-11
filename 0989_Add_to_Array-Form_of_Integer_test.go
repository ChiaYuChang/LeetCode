package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestAddToArrayFormOfInteger(t *testing.T) {
	type testCase struct {
		name string
		k    int
		num  []int
		ans  []int
	}

	tcs := []testCase{
		{
			name: "Simple_Case",
			k:    34,
			num:  []int{1, 2, 0, 0},
			ans:  []int{1, 2, 3, 4},
		},
		{
			name: "Case_w/_Carry",
			k:    181,
			num:  []int{2, 7, 4},
			ans:  []int{4, 5, 5},
		},
		{
			name: "Using ext array",
			k:    806,
			num:  []int{2, 1, 5},
			ans:  []int{1, 0, 2, 1},
		},
		{
			name: "Using more ext array",
			k:    110101,
			num:  []int{1, 0, 1, 0},
			ans:  []int{1, 1, 1, 1, 1, 1},
		},
		{
			name: "Final carry",
			k:    1,
			num:  []int{9, 9, 9},
			ans:  []int{1, 0, 0, 0},
		},
		{
			name: "Final carry",
			k:    1000,
			num:  []int{0},
			ans:  []int{1, 0, 0, 0},
		},
	}

	q := leetcode.Q0989{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				rst := q.AddToArrayForm(tc.num, tc.k)
				require.Equal(t, len(tc.ans), len(rst))
				for i := range tc.ans {
					require.Equal(t, tc.ans[i], rst[i])
				}
			},
		)
	}
}
