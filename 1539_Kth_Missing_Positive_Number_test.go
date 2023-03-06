package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindKthMissingPositive(t *testing.T) {
	type testCase struct {
		name string
		arr  []int
		k    int
		ans  int
	}

	tcs := []testCase{
		{
			name: "Case 1 General case",
			arr:  []int{2, 3, 4, 7, 11},
			k:    5,
			ans:  9,
		},
		{
			name: "Case 2 Greater than the last num",
			arr:  []int{1, 2, 3, 4},
			k:    2,
			ans:  6,
		},
		{
			name: "Case 3 More challenge general case",
			arr:  []int{1, 2, 3, 4, 5, 100},
			k:    73,
			ans:  78,
		},
		{
			name: "Case 4 Single element array",
			arr:  []int{2},
			k:    1,
			ans:  1,
		},
		{
			name: "Case 5 Two elements array",
			arr:  []int{3, 10},
			k:    2,
			ans:  2,
		},
	}

	q := leetcode.Q1539{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans,
					q.FindKthMissingPositive(tc.arr, tc.k))
			},
		)
	}
}
