package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestCarFleet(t *testing.T) {
	type testCase struct {
		name     string
		target   int
		position []int
		speed    []int
		ans      int
	}

	tcs := []testCase{
		{
			name:     "Case 1 general case",
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			ans:      3,
		},
		{
			name:     "Case 2 single car",
			target:   10,
			position: []int{3},
			speed:    []int{3},
			ans:      1,
		},
		{
			name:     "Case 3 merge to one",
			target:   100,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			ans:      1,
		},
		{
			name:     "Case 4 w/ res",
			target:   10,
			position: []int{6, 8},
			speed:    []int{3, 2},
			ans:      2,
		},
		{
			name:     "Case 5",
			target:   31,
			position: []int{5, 26, 18, 25, 29, 21, 22, 12, 19, 6},
			speed:    []int{7, 6, 6, 4, 3, 4, 9, 7, 6, 4},
			ans:      6,
		},
		{
			name:     "Case 6",
			target:   21,
			position: []int{1, 15, 6, 8, 18, 14, 16, 2, 19, 17, 3, 20, 5},
			speed:    []int{8, 5, 5, 7, 10, 10, 7, 9, 3, 4, 4, 10, 2},
			ans:      7,
		},
		{
			name:     "Case 7",
			target:   10,
			position: []int{4, 6},
			speed:    []int{3, 2},
			ans:      1,
		},
	}

	q := leetcode.Q0853{}
	for i := range tcs {
		tc := tcs[i]
		t.Log("==========", tc.name, "==========")
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(
					t, tc.ans,
					q.CarFleet(tc.target, tc.position, tc.speed),
				)
			},
		)
	}
}
