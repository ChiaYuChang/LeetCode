package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSuefaceArea(t *testing.T) {
	type testCase struct {
		name string
		grid [][]int
		sa   int
	}

	tcs := []testCase{
		{
			name: "Case 1",
			grid: [][]int{{1, 2}, {3, 4}},
			sa:   34,
		},
		{
			name: "Case 2",
			grid: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			sa:   32,
		},
		{
			name: "Case 3",
			grid: [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}},
			sa:   46,
		},
		{
			name: "Single cube",
			grid: [][]int{{10}},
			sa:   42,
		},
		{
			name: "Zero grids",
			grid: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			sa:   0,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.sa, leetcode.SurfaceArea(tc.grid))
		})
	}
}
